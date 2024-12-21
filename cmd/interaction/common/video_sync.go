package common

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"HuaTug.com/cmd/interaction/dal/db"
	"HuaTug.com/cmd/interaction/infras/client"
	"HuaTug.com/cmd/interaction/infras/redis"
	"HuaTug.com/cmd/model"
	"HuaTug.com/kitex_gen/videos"
	"HuaTug.com/pkg/constants"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/google/uuid"
)

// VideoSyncman 结构体
type VideoSyncman struct {
	ctx    context.Context
	cancel context.CancelFunc
}

// 新建 VideoSyncman 对象
func NewVideoSyncman() *VideoSyncman {
	ctx, cancel := context.WithCancel(context.Background())
	return &VideoSyncman{
		ctx:    ctx,
		cancel: cancel,
	}
}

type videoSyncData struct {
	vid          string
	likeList     *[]string
	dislikeList  *[]string
	visitCount   string
	commentCount string
}

// Run 方法，定期同步视频数据
// 首先进行一个初次同步，然后每隔一段时间进行一次同步
func (sm *VideoSyncman) Run() {
	if err := videoSyncMwWhenInit(); err != nil {
		panic(err)
	}

	go func() {
		for {
			time.Sleep(time.Second * 2) // 每隔 10 分钟执行一次
			select {
			case <-sm.ctx.Done():
				hlog.Info("Ok,停止同步[video]")
				return
			default:
			}
			if err := sm.syncVideoData(); err != nil {
				hlog.Error("同步视频数据时发生错误: ", err)
			}
		}
	}()
}

// 停止同步
func (sm *VideoSyncman) Stop() {
	sm.cancel()
}

// 同步视频数据的主逻辑
func (sm *VideoSyncman) syncVideoData() error {
	var (
		wg       sync.WaitGroup
		errChan  = make(chan error, 1)
		errChan2 = make(chan error, 1)
		isEnd    = false
		err      error
	)

	for i := 0; !isEnd; i++ {
		var vidList *[]string
		resp := new(videos.VideoIdListResponse)
		if resp, err = client.VideoIdList(context.Background(), &videos.VideoIdListRequest{
			PageNum:  int64(i),
			PageSize: 1000,
		}); err != nil {
			return fmt.Errorf("获取视频列表失败: %w", err)
		}
		isEnd = resp.IsEnd
		vidList = &resp.List

		for _, vid := range *vidList {
			videoData := newVideoSyncData(vid)
			videoId, _ := strconv.ParseInt(vid, 10, 64)
			wg.Add(4)
			// 并行获取数据 返回值全部都在访问函数的参数列表中
			go sm.getVisitCount(videoId, videoData, &wg, errChan)
			go sm.getCommentCount(videoId, videoData, &wg, errChan)
			go sm.getLikeList(videoId, videoData, &wg, errChan)
			go sm.getDislikeList(videoId, videoData, &wg, errChan)

			wg.Wait()

			select {
			case err := <-errChan:
				hlog.Error("处理视频数据时发生错误: ", err)
				continue
			default:
			}

			// 更新数据库和缓存
			wg.Add(2)
			go func() {
				defer wg.Done()
				if err := sm.updateVideoData(videoData); err != nil {
					hlog.Error("更新视频数据失败: ", err)
					errChan2 <- err
				}
			}()
			go func() {
				defer wg.Done()
				if err := sm.updateLikeCount(videoData); err != nil {
					hlog.Error("更新视频点赞信息失败: ", err)
					errChan2 <- err
				}
			}()
			wg.Wait()

			select {
			case err := <-errChan2:
				hlog.Error("处理视频数据时发生错误: ", err)
			default:
				continue
			}
		}
	}
	return nil
}

// 获取视频访问次数
func (sm *VideoSyncman) getVisitCount(vid int64, data *videoSyncData, wg *sync.WaitGroup, errChan chan error) {
	defer wg.Done()
	resp, err := client.GetVideoVisitCountInRedis(context.Background(), &videos.GetVideoVisitCountInRedisRequest{
		VideoId: vid,
	})

	if err != nil {
		errChan <- fmt.Errorf("获取视频访问次数失败: %w", err)
		return
	}

	hlog.Info(fmt.Sprintf("Response: %+v", resp))
	if resp == nil {
		return
	}
	data.visitCount = fmt.Sprint(resp.VisitCount)
}

// 获取视频评论数
func (sm *VideoSyncman) getCommentCount(vid int64, data *videoSyncData, wg *sync.WaitGroup, errChan chan error) {
	defer wg.Done()
	commentCount, err := db.GetVideoCommentCount(sm.ctx, vid)
	if err != nil {
		errChan <- fmt.Errorf("获取视频评论数失败: %w", err)
		return
	}
	data.commentCount = fmt.Sprint(commentCount)
}

// 获取点赞列表
func (sm *VideoSyncman) getLikeList(vid int64, data *videoSyncData, wg *sync.WaitGroup, errChan chan error) {
	defer wg.Done()
	likeList, err := redis.GetNewUpdateVideoLikeList(vid)
	if err != nil {
		errChan <- fmt.Errorf("获取点赞列表失败: %w", err)
		return
	}
	data.likeList = likeList
}

// 获取取消点赞列表
func (sm *VideoSyncman) getDislikeList(vid int64, data *videoSyncData, wg *sync.WaitGroup, errChan chan error) {
	defer wg.Done()
	dislikeList, err := redis.GetNewDeleteVideoLikeList(vid)
	if err != nil {
		errChan <- fmt.Errorf("获取取消点赞列表失败: %w", err)
		return
	}
	data.dislikeList = dislikeList
}

// 更新视频数据到数据库、缓存和 Elasticsearch
func (sm *VideoSyncman) updateVideoData(data *videoSyncData) error {
	// 更新点赞信息
	for _, uid := range *data.likeList {
		if err := sm.updateLikeInfo(data.vid, uid); err != nil {
			return fmt.Errorf("更新点赞信息失败: %w", err)
		}
	}

	// 更新取消点赞信息
	for _, uid := range *data.dislikeList {
		if err := sm.updateDislikeInfo(data.vid, uid); err != nil {
			return fmt.Errorf("更新取消点赞信息失败: %w", err)
		}
	}

	if err := sm.updateCommentCount(data); err != nil {
		return fmt.Errorf("更新评论数失败: %w", err)
	}
	// 更新访问次数
	// if err := sm.updateVisitCount(data); err != nil {
	// 	return fmt.Errorf("更新访问次数失败: %w", err)
	// }

	// 更新 Elasticsearch 数据
	// if err := sm.updateElasticSearch(data); err != nil {
	// 	return fmt.Errorf("更新 Elasticsearch 数据失败: %w", err)
	// }

	return nil
}

// 更新点赞信息 （在进行同步时，将点赞信息从动态空间移动到静态空间）
func (sm *VideoSyncman) updateLikeInfo(vid, uid string) error {
	videoId, _ := strconv.ParseInt(vid, 10, 64)
	userId, _ := strconv.ParseInt(uid, 10, 64)
	videolikeId := uuid.New().ID()
	if err := db.CreateVideoLike(sm.ctx, &model.VideoLike{
		VideoLikesId: int64(videolikeId),
		UserId:       userId,
		VideoId:      videoId,
		CreatedAt:    time.Now().Format(constants.DataFormate),
		DeletedAt:    "",
	}); err != nil {
		// 如果创建失败，尝试删除动态空间中的数据
		go redis.DeleteVideoLikeInfoFromDynamicSpace(videoId, userId)
		return err
	}

	// 实现将点赞信息同步到静态空间，并删除动态空间中的数据
	if err := redis.AppendVideoLikeInfoToStaticSpace(videoId, userId); err != nil {
		return err
	}
	if err := redis.DeleteVideoLikeInfoFromDynamicSpace(videoId, userId); err != nil {
		return err
	}
	return nil
}

// 更新取消点赞信息
func (sm *VideoSyncman) updateDislikeInfo(vid, uid string) error {
	videoId, _ := strconv.ParseInt(vid, 10, 64)
	userId, _ := strconv.ParseInt(uid, 10, 64)
	// 先执行数据库操作 后执行缓存操作
	if err := db.DeleteVideoLike(context.Background(), videoId, userId); err != nil {
		return err
	}
	if err := redis.DeleteVideoLikeInfoFromDynamicSpace(videoId, userId); err != nil {
		return err
	}
	return nil
}

// 更新视频评论数
func (sm *VideoSyncman) updateCommentCount(data *videoSyncData) error {
	videoId, _ := strconv.ParseInt(data.vid, 10, 64)
	commentCount, err := strconv.ParseInt(data.commentCount, 10, 64)
	if err != nil {
		return err
	}
	if _, err = client.UpdateVideoCommentCount(sm.ctx, &videos.UpdateVideoCommentCountRequest{VideoId: videoId, CommentCount: commentCount}); err != nil {
		hlog.Info(err)
		return err
	}
	return nil
}

// 更新视频点赞量
func (sm *VideoSyncman) updateLikeCount(data *videoSyncData) error {
	videoId, _ := strconv.ParseInt(data.vid, 10, 64)
	his_count, err := db.GetVideoLikeList(sm.ctx, videoId)
	if err != nil {
		hlog.Info(err)
		return err
	}

	fin_likecount := len(*his_count) + len(*data.likeList)
	_, err = client.UpdateVideoHisLikeCount(sm.ctx, &videos.UpdateVideoHisLikeCountRequest{
		VideoId:      videoId,
		HisLikeCount: int64(len(*his_count)),
	})
	if err != nil {
		hlog.Info(err)
		return err
	}

	_, err = client.UpdateVideoLikeCount(sm.ctx, &videos.UpdateLikeCountRequest{
		VideoId:   videoId,
		LikeCount: int64(fin_likecount),
	})
	if err != nil {
		hlog.Info(err)
		return err
	}
	return nil
}

// 更新访问次数
func (sm *VideoSyncman) updateVisitCount(data *videoSyncData) error {
	hlog.Info(data.visitCount)
	visitCount, err := strconv.ParseInt(data.visitCount, 10, 64)
	if err != nil {
		hlog.Info(err)
		return err
	}
	hlog.Info(visitCount)

	videoId, _ := strconv.ParseInt(data.vid, 10, 64)
	_, err = client.UpdateVisitCount(context.Background(), &videos.UpdateVisitCountRequest{
		VideoId:    videoId,
		VisitCount: visitCount,
	})
	if err != nil {
		hlog.Info(err)
		return err
	}
	//ToDo :是否考虑使用RabbitMQ 来进行通信 降低通信的压力
	time.Sleep(10 * time.Second)
	return nil
}

// 更新 Elasticsearch 数据
// func (sm *VideoSyncman) updateElasticSearch(data *videoSyncData) error {
// 	likeCount := len(*data.likeList)
// 	return elasticsearch.UpdateVideoLikeVisitAndCommentCount(data.vid, fmt.Sprint(likeCount), data.visitCount, data.commentCount)
// }

// 初始化时同步视频数据
func videoSyncMwWhenInit() error {
	var (
		isEnd = false
		list  *[]string
		err   error
	)

	for i := 0; !isEnd; i++ {
		resp := new(videos.VideoIdListResponse)
		resp, err = client.VideoIdList(context.Background(), &videos.VideoIdListRequest{
			PageSize: 1000,
			PageNum:  int64(i),
		})
		if err != nil {
			return fmt.Errorf("获取视频列表失败: %w", err)
		}
		if resp == nil {
			hlog.Info("is nil")
			return errors.New("Failed to get VideoIdList")
		}
		isEnd = resp.IsEnd
		list = &resp.List
		var (
			wg       sync.WaitGroup
			errChan  = make(chan error, 1)
			syncList = make([]videoSyncData, 0)
		)

		for _, vid := range *list {
			videoId, _ := strconv.ParseInt(vid, 10, 64)
			data := newVideoSyncData(vid)
			wg.Add(2)
			go NewVideoSyncman().getLikeList(videoId, data, &wg, errChan)
			//go NewVideoSyncman().getVisitCount(videoId, data, &wg, errChan)
			go NewVideoSyncman().getCommentCount(videoId, data, &wg, errChan)
			wg.Wait()

			select {
			case err := <-errChan:
				return err
			default:
			}

			syncList = append(syncList, *data)
		}

		errChan = make(chan error, 1)
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := videoSyncDB2Redis(&syncList); err != nil {
				errChan <- err
			}
		}()
		wg.Wait()

		select {
		case err := <-errChan:
			return err
		default:
		}
	}

	return nil
}

// 工具函数：创建视频同步数据对象
func newVideoSyncData(vid string) *videoSyncData {
	return &videoSyncData{
		vid: vid,
	}
}

func videoSyncDB2Redis(syncList *[]videoSyncData) error {
	for _, item := range *syncList {
		videoId, _ := strconv.ParseInt(item.vid, 10, 64)
		if err := redis.PutVideoLikeInfo(videoId, item.likeList); err != nil {
			return err
		}
	}
	return nil
}
