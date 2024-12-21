package service

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"HuaTug.com/cmd/model"
	"HuaTug.com/cmd/video/dal/db"
	"HuaTug.com/cmd/video/infras/client"
	"HuaTug.com/cmd/video/infras/redis"
	"HuaTug.com/kitex_gen/base"
	"HuaTug.com/kitex_gen/users"
	"HuaTug.com/kitex_gen/videos"
	"HuaTug.com/pkg/constants"
	"HuaTug.com/pkg/errno"
	"HuaTug.com/pkg/oss"
	"HuaTug.com/pkg/utils"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/pkg/errors"
)

type VideoUploadService struct {
	ctx context.Context
}

func NewVideoUploadService(ctx context.Context) *VideoUploadService {
	return &VideoUploadService{
		ctx: ctx,
	}
}

var (
	TempVideoFolderPath string
)

func (service *VideoUploadService) NewCancleUploadEvent(req *videos.VideoPublishCancleRequest) error {
	if req.Uuid == `` {
		return errno.RequestErr
	}
	if err := service.deleteTempDir(fmt.Sprint(req.UserId) + `_` + req.Uuid); err != nil {
		return errors.WithMessage(err, "Failed to deleteTemDir")
	}
	if err := redis.DeleteVideoEvent(service.ctx, req.Uuid, fmt.Sprint(req.UserId)); err != nil {
		return errors.WithMessage(err, "Failed to DeleteVideo")
	}
	return nil
}

// 合并分片为一个完整的视频文件
func MergeChunks(chunkDir, OutputFile, filename string, totalChunks int) error {
	// 创建输出文件
	// 获取 OutputFile 的目录部分
	outputDir := filepath.Dir(OutputFile)

	// 创建目录（如果不存在）
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		hlog.Error("Failed to create directories:", err)
		return errors.WithMessage(err, "Failed to create directories: ")
	}
	output, err := os.Create(OutputFile)
	if err != nil {
		return errors.WithMessage(err, "Failed to create output file: ")
	}
	defer output.Close()

	// 按照顺序读取并写入每个分片
	for i := 1; i <= totalChunks; i++ {
		chunkPath := filepath.Join(chunkDir, fmt.Sprintf("/videos_part_%d", i))
		chunkFile, err := os.Open(chunkPath)
		if err != nil {
			return fmt.Errorf("Failed to open chunk %d: %v", i, err)
		}
		// 将当前分片写入到输出文件
		_, err = io.Copy(output, chunkFile)
		if err != nil {
			chunkFile.Close()
			return fmt.Errorf("Failed to write chunk %d to output: %v", i, err)
		}
		chunkFile.Close()
	}
	return nil
}

func (service *VideoUploadService) CompleteUpload(uuid, uid, vid string, chunkDir string, totalChunks int) (string, error) {
	// 合并分片
	OutputFile := filepath.Join(chunkDir, "merged_video.mp4")
	temp, err := redis.GetChunkInfo(uid, uuid)
	if err != nil {
		return "", fmt.Errorf("Failed to get chunk info: %v", err)
	}

	var wg sync.WaitGroup
	errs := make(chan error, totalChunks)

	for i := 1; i <= totalChunks; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			chunkPath := filepath.Join(chunkDir, fmt.Sprintf("/videos_part_%d", i))
			// 上传分片逻辑
			if _, err := oss.UploadVideo(chunkPath, vid); err != nil {
				errs <- fmt.Errorf("Failed to upload chunk %d: %v", i, err)
			}
		}(i)
	}

	wg.Wait()
	close(errs)

	// 检查上传过程中是否有错误
	for err := range errs {
		if err != nil {
			return "", err
		}
	}

	if err := MergeChunks(chunkDir, OutputFile, temp[1], totalChunks); err != nil {
		return "", errors.WithMessage(err, "Failed to merge chunks: ")
	}
	return OutputFile, nil
}
func (service *VideoUploadService) NewUploadCompleteEvent(req *videos.VideoPublishCompleteRequest) error {
	if req.Uuid == `` {
		return errno.RequestErr
	}
	reallyComplete, err := redis.IsChunkAllRecorded(service.ctx, req.Uuid, fmt.Sprint(req.UserId))
	if err != nil {
		return errno.RedisErr
	}
	if !reallyComplete {
		return errors.New("Not all chunks have been uploaded")
	}

	// m3u8name, err := redis.GetM3U8Filename(service.ctx, req.Uuid, fmt.Sprint(req.UserId))
	// if err != nil {
	// 	return errno.RedisErr
	// }

	// err = utils.M3u8ToMp4(TempVideoFolderPath+fmt.Sprint(req.UserId)+`_`+req.Uuid+`/`+m3u8name,
	// 	TempVideoFolderPath+fmt.Sprint(req.UserId)+`_`+req.Uuid+`/`+`video.mp4`)
	// if err != nil {
	// 	return errno.ServiceErr
	// }

	// err = utils.GenerateMp4CoverJpg(TempVideoFolderPath+fmt.Sprint(req.UserId)+`_`+req.Uuid+`/`+`video.mp4`,
	// 	TempVideoFolderPath+fmt.Sprint(req.UserId)+`_`+`cover.jpg`)
	// if err != nil {
	// 	return errno.ServiceErr
	// }
	var Chunkdir = fmt.Sprint(req.UserId) + "_" + fmt.Sprint(req.Uuid)
	info, err := redis.GetChunkInfo(fmt.Sprint(req.UserId), req.Uuid)
	if err != nil {
		return errno.RedisErr
	}

	if err != nil {
		return errno.ServiceErr
	}
	TotalNumber, _ := strconv.ParseInt(info[0], 10, 64)
	vid, err := db.GetMaxVideoId(service.ctx)
	if err != nil {
		return errors.New("Failed to get max video_id")
	}
	videofile, err := service.CompleteUpload(req.Uuid, fmt.Sprint(req.UserId), vid, Chunkdir, int(TotalNumber))
	if err != nil {
		hlog.Info("Error:", err)
		return errors.WithMessage(err, "Failed to upload the file")
	}
	var (
		videoUrl, coverUrl string
		resp               *users.GetUserInfoResponse
		wg                 sync.WaitGroup
	)
	errChan := make(chan error, 3)

	wg.Add(3)
	go func() {
		videoUrl, err = oss.UploadVideo(videofile, vid)
		if err != nil {
			errChan <- errno.OssErr
		}
		wg.Done()
	}()
	go func() {
		//temp := `/home/xuzh/Golang/Work-6/cmd/video/` + fmt.Sprint(req.UserId) + `_` + req.Uuid
		temp := `../` + fmt.Sprint(req.UserId) + `_` + req.Uuid
		coverPath, err := utils.GetVideoThumnail(videofile, temp)
		if err != nil {
			errChan <- errors.WithMessage(err, "Failed to GetVideoThumnail to minio")
		}
		coverUrl, err = oss.UploadVideoCover(coverPath, vid)
		if err != nil {
			errChan <- errors.WithMessage(err, "Failed to UploadVideoCover to minio")
		}
		wg.Done()
	}()
	// ToDo: 用于获取用户信息 完善ElasticSearch索引
	go func() {
		resp = new(users.GetUserInfoResponse)
		resp, err = client.UserClient.GetUserInfo(service.ctx, &users.GetUserInfoRequest{
			UserId: req.UserId,
		})
		if err != nil {
			errChan <- err
		}
		hlog.Info(resp)
		wg.Done()
	}()
	wg.Wait()
	select {
	case err := <-errChan:
		if err != nil {
			hlog.Info(err)
			return err
		}
	default:
		break
	}
	d := base.Video{
		Title:       info[1],
		Description: info[2],
		UserId:      req.UserId,
		VisitCount:  0,
		LabelNames:  info[3],
		Category:    info[4],
		CreatedAt:   time.Now().Format(constants.DataFormate),
		UpdatedAt:   time.Now().Format(constants.DataFormate),
		VideoUrl:    videoUrl,
		CoverUrl:    coverUrl,
		DeletedAt:   "0",
	}
	err = db.InsertVideo(service.ctx, &d)
	if err != nil {
		return errno.MysqlErr
	}

	err = redis.DeleteVideoEvent(service.ctx, req.Uuid, fmt.Sprint(req.UserId))
	if err != nil {
		return errno.ServiceErr
	}

	err = service.deleteTempDir(fmt.Sprint(req.UserId) + `_` + req.Uuid)
	if err != nil {
		return errno.ServiceErr
	}
	return nil
}

func (service *VideoUploadService) NewUploadingEvent(req *videos.VideoPublishUploadingRequest) error {
	if req.Filename == `` || req.Uuid == `` || req.ChunkNumber <= 0 {
		return errno.RequestErr
	}
	data := req.Data

	if !service.isMD5Same(data, req.Md5) {
		return errors.New("Data proccess failed")
	}
	if req.IsM3u8 {
		err := redis.RecordM3U8Filename(service.ctx, req.Uuid, fmt.Sprint(req.UserId), req.Filename)
		if err != nil {
			return errors.WithMessage(err, "RecordM3U8Filename failed!")
		}
	}

	hlog.Info(TempVideoFolderPath)
	err := service.saveTempDate(TempVideoFolderPath+fmt.Sprint(req.UserId)+`_`+req.Uuid+`/`+req.Filename, data)
	if err != nil {
		return errors.WithMessage(err, "SaveTempDate failed!")
	}
	err = redis.DoneChunkEvent(service.ctx, req.Uuid, fmt.Sprint(req.UserId), req.ChunkNumber)
	if err != nil {
		return errors.WithMessage(err, "DoneChunkEvent failed!")
	}
	return nil
}

func (service *VideoUploadService) NewUploadEvent(req *videos.VideoPublishStartRequest) (string, error) {
	var (
		uuid = ``
		uid  = fmt.Sprint(req.UserId)
		err  error
	)
	if req.Title == `` || req.ChunkTotalNumber <= 0 {
		return ``, errno.RequestErr
	}
	uuid, err = redis.NewVideoEvent(service.ctx, req.Title, req.Description, uid, fmt.Sprint(req.ChunkTotalNumber), req.LabName, req.Category)
	if err != nil {
		return ``, errno.RedisErr
	}
	cwd, _ := os.Getwd()
	hlog.Info(cwd)
	//os.Mkdir在创建目录时 是在当前代码的工作目录下进行，即main函数的位置
	if err := os.Mkdir(uid+`_`+uuid, os.ModePerm); err != nil {
		if err := redis.DeleteVideoEvent(service.ctx, uuid, uid); err != nil {
			return ``, errno.RedisErr
		}
		return ``, errno.ServiceErr
	}
	return uuid, nil
}

func (service *VideoUploadService) NewDeleteEvent(req *videos.VideoDeleteRequest) error {
	var (
		wg      sync.WaitGroup
		errChan = make(chan error, 3)
	)
	//在这里进行删除操作时 首先删除数据库中的数据 然后删除与这条视频相关的所有数据（例如interaction模块中的评论）
	wg.Add(3)
	go func() {
		if err := db.DeleteVideo(service.ctx, fmt.Sprint(req.VideoId), fmt.Sprint(req.UserId)); err != nil {
			errChan <- errors.WithMessage(err, "Failed to DeleteVideo")
		}
		wg.Done()
	}()
	go func() {

		wg.Done()
	}()
	go func() {

		wg.Done()
	}()
	wg.Wait()
	select {
	case err := <-errChan:
		return err
	default:
		break
	}
	return nil
}
func (service *VideoUploadService) NewSearchEvent(req *videos.VideoSearchRequest) (*videos.VideoSearchResponse, error) {
	return nil, nil
}
func (service *VideoUploadService) NewIdListEvent(req *videos.VideoIdListRequest) (bool, *[]string, error) {
	list, err := db.GetVideoIdList(service.ctx, req.PageNum, req.PageSize)
	if err != nil {
		return true, nil, errors.WithMessage(err, "Failed to get list by id")
	}
	return len(*list) < int(req.PageSize), list, nil
}

func (service *VideoUploadService) NewUpdateVideoVisitCountEvent(req *videos.UpdateVisitCountRequest) error {
	err := db.UpdateVideoVisit(service.ctx, req.VideoId, req.VisitCount)
	if err != nil {
		return errors.WithMessage(err, "Failed to update visitcount")
	}
	return nil
}

func (service *VideoUploadService) NewUpdateVideoCommentCountEvent(req *videos.UpdateVideoCommentCountRequest) error {
	err := db.UpdateVideoCommentCount(service.ctx, req.VideoId, req.CommentCount)
	if err != nil {
		return errors.WithMessage(err, "Failed to update commentcount")
	}
	return nil
}

func (service *VideoUploadService) NewUpdateVideoLikeCountEvent(req *videos.UpdateLikeCountRequest) error {
	err := db.UpdateVideoLikeCount(service.ctx, req.VideoId, req.LikeCount)
	if err != nil {
		return errors.WithMessage(err, "Failed to update likecount")
	}
	return nil
}

func (service *VideoUploadService) NewUpdateVideoHisLikeCountEvent(req *videos.UpdateVideoHisLikeCountRequest) error {
	err := db.UpdateVideoHisLikeCount(service.ctx, req.VideoId, req.HisLikeCount)
	if err != nil {
		return errors.WithMessage(err, "Failed to update likecount")
	}
	return nil
}
func (service *VideoUploadService) NewGetVisitCountEvent(req *videos.GetVideoVisitCountRequest) (count int64, err error) {
	count, err = db.GetVideoVisitCount(service.ctx, fmt.Sprint(req.VideoId))
	if err != nil {
		hlog.Info(err)
		return -1, errors.WithMessage(err, "Failed to get VideoVisitCount")
	}
	hlog.Info(count)
	return count, nil
}

func (service *VideoUploadService) NewVideoVisitEvent(req *videos.VideoVisitRequest) (*base.Video, error) {
	var (
		wg      sync.WaitGroup
		errChan = make(chan error, 2)
		data    *base.Video
		err     error
	)
	wg.Add(4)
	go func() {
		data, err = db.GetVideo(service.ctx, req.VideoId)
		if err != nil {
			errChan <- errors.WithMessage(err, "Get VideoInfo failed")
		}
		wg.Done()
	}()
	go func() {
		if err := redis.IncrVideoVisitInfo(fmt.Sprint(req.VideoId)); err != nil {
			errChan <- errors.WithMessage(err, "Failed to NewVideoVisitEvent")
		}
		wg.Done()
	}()
	go func() {
		watch := &model.UserVideoWatchHistory{
			UserId:    req.FromId,
			VideoId:   req.VideoId,
			WatchTime: time.Now().Format(constants.DataFormate),
		}
		if err := db.AddUserVideoWatchHistory(service.ctx, watch); err != nil {
			errChan <- errors.WithMessage(err, "Failed to AddUserVideoWatchHistory")
		}
		wg.Done()
	}()
	go func() {
		view := &model.UserBehavior{
			UserId:       req.FromId,
			VideoId:      req.VideoId,
			BehaviorType: "view",
			BehaviorTime: time.Now().Format(constants.DataFormate),
		}
		if err := db.AddUserViewBehavior(service.ctx, view); err != nil {
			errChan <- errors.WithMessage(err, "Failed to AddUserViewBehavior")
		}
		wg.Done()
	}()
	wg.Wait()

	select {
	case result := <-errChan:
		return nil, result
	default:
		break
	}
	return data, nil
}
func (service *VideoUploadService) NewGetVisitCountInRedisEvent(req *videos.GetVideoVisitCountInRedisRequest) (int64, error) {
	data, err := redis.GetVideoVisitCount(fmt.Sprint(req.VideoId))
	if err != nil {
		return -1, errors.WithMessage(err, "Failed to get visitcount")
	}
	return data, nil
}
func (service *VideoUploadService) deleteTempDir(path string) error {
	return os.RemoveAll(path)
}

func (service *VideoUploadService) saveTempDate(path string, data []byte) error {
	return os.WriteFile(path, data, 0777)
}

func (service *VideoUploadService) isMD5Same(data []byte, md5 string) bool {
	return utils.GetBytesMD5(data) == md5
}
