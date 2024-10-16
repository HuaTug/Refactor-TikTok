package db

import (
	"context"
	"fmt"

	"sync"

	"HuaTug.com/cmd/model"
	"HuaTug.com/kitex_gen/base"
	"HuaTug.com/kitex_gen/videos"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func Feedlist(ctx context.Context, req *videos.FeedServiceRequest) ([]*model.Video, error) {
	var video []*model.Video
	if err := DB.WithContext(ctx).Model(&model.Video{}).Where("created_at<?", req.LastTime).Find(&video); err != nil {
		return video, errors.Wrapf(err.Error, "FeedList failed,err:%v", err)
	}
	return video, nil
}

func Videolist(ctx context.Context, req *videos.VideoFeedListRequest) ([]*base.Video, int64, error) {
	var video []*base.Video
	var count int64
	if err := DB.WithContext(ctx).Model(&base.Video{}).Where("user_id=?", req.UserId).Count(&count).Limit(int(req.PageSize)).
		Offset(int((req.PageNum - 1) * req.PageSize)).Find(&video); err != nil {
		logrus.Info(err)
		return video, count, errors.Wrapf(err.Error, "VideoList failed,err:%v", err)
	}
	return video, count, nil
}

func Videosearch(ctx context.Context, req *videos.VideoSearchRequest) ([]*base.Video, int64, error) {
	var wg sync.WaitGroup
	var video2 []*base.Video
	var count int64
	var err error
	if req.Keyword != "" {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err = DB.WithContext(ctx).Model(&base.Video{}).
				Where("title like ? And created_at<? And created_at>?", "%"+req.Keyword+"%", req.ToDate, req.FromDate).
				Count(&count).
				Limit(int(req.PageSize)).Offset(int((req.PageNum - 1) * req.PageSize)).
				Find(&video2).Error
		}()
		if err != nil {
			return video2, count, errors.Wrapf(err, "VideoSearch failed,err:%v", err)
		}
		wg.Wait()
	}
	return video2, count, nil
}

func FindVideo(ctx context.Context, videoId int64) (video *base.Video, err error) {
	if err := DB.WithContext(ctx).Model(&base.Video{}).Where("video_id=?", videoId).Find(&video); err != nil {
		return video, errors.Wrapf(err.Error, "FindVideo failed,err:%v", err)
	}
	return video, nil
}

func InsertVideo(ctx context.Context, video *model.Video) (string, error) {
	if err := DB.WithContext(ctx).Create(video).Error; err != nil {
		return ``, err
	} else {
		var maxId int64
		DB.WithContext(ctx).Model(&model.Video{}).Select("MAX(video_id)").Scan(&maxId)
		return fmt.Sprint(maxId), nil
	}
}

func GetVideo(ctx context.Context, vid int64) (*base.Video, error) {
	var data base.Video
	if err := DB.WithContext(ctx).Model(&base.Video{}).Where("video_id = ?", vid).Find(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func UpdateVideoUrl(ctx context.Context, videoUrl, coverUrl, vid string) error {
	if err := DB.WithContext(ctx).Model(&base.Video{}).Where("video_id = ?", vid).Update("video_url", videoUrl).Error; err != nil {
		return err
	}
	if err := DB.WithContext(ctx).Model(&base.Video{}).Where("video_id = ?", vid).Update("cover_url", coverUrl).Error; err != nil {
		return err
	}
	return nil
}

func UpdateVideoVisit(ctx context.Context, vid, visitCount int64) error {
	if err := DB.WithContext(ctx).Model(&base.Video{}).Where("video_id = ?", vid).Update("visit_count", visitCount).Error; err != nil {
		return err
	}
	return nil
}

func DeleteVideo(ctx context.Context, vid, uid string) error {
	result := DB.Model(&base.Video{}).Where("video_id = ? And user_id=? ", vid, uid).Delete(&base.Video{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("No rows has been affected")
	}
	return nil
}

func GetVideoVisitCount(ctx context.Context, vid string) (count int64, err error) {
	//Scan用于将查询结果集映射到某一个值上 Scan和Count的区别使用·
	if err = DB.Model(&base.Video{}).Select("visit_count").Where("video_id = ?", vid).Scan(&count).Error; err != nil {
		return 0, err
	}
	return count, err

}

func GetVideoIdList(ctx context.Context, pageNum, pageSize int64) (*[]string, error) {
	list := make([]string, 0)
	if err := DB.Model(&base.Video{}).Select("video_id").Limit(int(pageSize)).Offset(int(pageNum)).Scan(&list).Error; err != nil {
		hlog.Info(err)
		return nil, err
	}
	return &list, nil
}

func GetVideoInfo(ctx context.Context, videoId int64) (*base.Video, error) {
	video := new(base.Video)
	var err error
	if err = DB.WithContext(ctx).Model(&base.Video{}).Where("video_id = ?", videoId).Find(video).Error; err != nil {
		return nil, errors.WithMessage(err, "Failed to get VideoInfo")
	}
	return video, nil
}
