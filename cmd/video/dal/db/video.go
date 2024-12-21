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

func Feedlist(ctx context.Context, req *videos.FeedServiceRequest) ([]*base.Video, error) {
	var video []*base.Video
	if err := DB.WithContext(ctx).Model(&base.Video{}).Where("created_at<?", req.LastTime).Find(&video); err != nil {
		return video, errors.Wrapf(err.Error, "FeedList failed,err:%v", err)
	}
	return video, nil
}

// 获取用户发布的视频
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

func InsertVideo(ctx context.Context, video *base.Video) error {
	if err := DB.WithContext(ctx).Create(video).Error; err != nil {
		return err
	}
	return nil
}
func GetMaxVideoId(ctx context.Context) (string, error) {
	var maxId *int64
	if err := DB.WithContext(ctx).Model(&base.Video{}).Select("MAX(video_id)").Scan(&maxId).Error; err != nil {
		return "", err
	}
	if maxId == nil {
		return "1", nil
	}

	return fmt.Sprint(*maxId), nil
}
func GetVideo(ctx context.Context, vid int64) (*base.Video, error) {
	var data base.Video
	if err := DB.WithContext(ctx).Model(&base.Video{}).Where("video_id = ?", vid).Find(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

// 对于视频列表的查询
func GetVideoByVideoId(ctx context.Context, vid []int64) ([]*base.Video, error) {
	var data []*base.Video
	if err := DB.WithContext(ctx).Model(&base.Video{}).Where("video_id IN (?)", vid).Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
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

func UpdateVideoCommentCount(ctx context.Context, vid, commentCount int64) error {
	if err := DB.WithContext(ctx).Model(&base.Video{}).Where("video_id =?", vid).Update("comment_count", commentCount).Error; err != nil {
		return err
	}
	return nil
}

func UpdateVideoLikeCount(ctx context.Context, vid, likeCount int64) error {
	if err := DB.WithContext(ctx).Model(&base.Video{}).Where("video_id =?", vid).Update("likes_count", likeCount).Error; err != nil {
		return err
	}
	return nil
}

func UpdateVideoHisLikeCount(ctx context.Context, vid, hisLikeCount int64) error {
	if err := DB.WithContext(ctx).Model(&base.Video{}).Where("video_id =?", vid).Update("history_count", hisLikeCount).Error; err != nil {
		return err
	}
	return nil
}

func UpdateVideoShareCount(ctx context.Context, vid, shareCount int64) error {
	if err := DB.WithContext(ctx).Model(&base.Video{}).Where("video_id =?", vid).Update("share_count", shareCount).Error; err != nil {
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

func GetVideoShareCount(ctx context.Context, vid string) (count int64, err error) {
	if err := DB.WithContext(ctx).Model(&model.VideoShare{}).Where("video_id = ?", vid).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, err
}

func GetVideoIdList(ctx context.Context, pageNum, pageSize int64) (*[]string, error) {
	list := make([]string, 0)
	if err := DB.Model(&base.Video{}).Select("video_id").Offset(int(pageNum-1) * int(pageSize)).Limit(int(pageSize)).Scan(&list).Error; err != nil {
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

func CreateFavorite(ctx context.Context, fav *base.Favorite) error {
	if err := DB.WithContext(ctx).Model(&base.Favorite{}).Create(fav).Error; err != nil {
		return errors.WithMessage(err, "Failed to create Favorite")
	}
	return nil
}

// 获取用户收藏列表（有多少个收藏夹）
func GetFavoriteList(ctx context.Context, req *videos.GetFavoriteListRequest) ([]*base.Favorite, error) {
	var favList []*base.Favorite
	hlog.Info(req.UserId)
	if err := DB.WithContext(ctx).Model(&base.Favorite{}).Where("user_id=?", req.UserId).
		Offset((int(req.PageNum) - 1) * int(req.PageSize)).Limit(int(req.PageSize)).Find(&favList).Error; err != nil {
		return nil, errors.WithMessage(err, "Failed to get FavoriteList")
	}
	return favList, nil
}

func AddVideoToFavorite(ctx context.Context, fav_vid *model.FavoritesVideos) error {
	if err := DB.WithContext(ctx).Model(&model.FavoritesVideos{}).Create(fav_vid).Error; err != nil {
		return errors.WithMessage(err, "Failed to add VideoToFavorite")
	}
	return nil
}

func GetVideoIdFromFavorite(ctx context.Context, user_id, favorite_id int64) ([]int64, error) {
	var videoIds []int64
	if err := DB.WithContext(ctx).Model(&model.FavoritesVideos{}).Where("user_id = ? and favorite_id = ?", user_id, favorite_id).Select("video_id").Scan(&videoIds).Error; err != nil {
		return videoIds, errors.WithMessage(err, "Failed to get VideoFromList")
	}
	hlog.Info(videoIds)
	return videoIds, nil
}

// 从视频收藏中获取视频列表
func GetFavoriteVideoList(ctx context.Context, req *videos.GetFavoriteVideoListRequest) ([]*base.Video, error) {
	var video []*base.Video
	videoIds, err := GetVideoIdFromFavorite(ctx, req.UserId, req.FavoriteId)
	if err != nil {
		return video, err
	}
	if len(videoIds) == 0 {
		return video, nil
	}

	if err := DB.WithContext(ctx).Model(&base.Video{}).Where("video_id in?", videoIds).Find(&video).Error; err != nil {
		return video, errors.WithMessage(err, "Failed to get VideoFromList")
	}
	return video, nil
}

func GetVideoFromFavorite(ctx context.Context, req *videos.GetVideoFromFavoriteRequest) (*base.Video, error) {
	var video *base.Video
	if err := DB.WithContext(ctx).Model(&base.Video{}).Where("user_id = ? and video_id = ?", req.UserId, req.VideoId).Find(&video).Error; err != nil {
		return nil, errors.WithMessage(err, "Failed to get VideoFromFavorite")
	}
	return video, nil
}

func DeleteFavorite(ctx context.Context, req *videos.DeleteFavoriteRequest) error {
	go DeleteAllVideoFromFavorite(ctx, req.UserId, req.FavoriteId)
	if err := DB.WithContext(ctx).Model(&model.Favorite{}).Where("user_id =? and favorite_id =?", req.UserId, req.FavoriteId).Delete(&base.Favorite{}).Error; err != nil {
		return errors.WithMessage(err, "Failed to delete Favorite")
	}
	return nil
}

func DeleteVideoFromFavorite(ctx context.Context, req *videos.DeleteVideoFromFavoriteRequest) error {
	if err := DB.WithContext(ctx).Model(&model.FavoritesVideos{}).Where("user_id =? and video_id =?", req.UserId, req.VideoId).Delete(&model.FavoritesVideos{}).Error; err != nil {
		return errors.WithMessage(err, "Failed to delete VideoFromFavorite")
	}
	return nil
}

func DeleteAllVideoFromFavorite(ctx context.Context, user_id, favorite_id int64) error {
	if err := DB.WithContext(ctx).Model(&model.FavoritesVideos{}).Where("user_id =? and favorite_id =?", user_id, favorite_id).Delete(&model.FavoritesVideos{}).Error; err != nil {
		return errors.WithMessage(err, "Failed to delete VideoFromFavorite")
	}
	return nil
}

func SharedVideo(ctx context.Context, share *model.VideoShare) error {
	if err := DB.WithContext(ctx).Model(&model.VideoShare{}).Create(share).Error; err != nil {
		return errors.WithMessage(err, "Failed to shared Video")
	}
	return nil
}

func AddUserVideoWatchHistory(ctx context.Context, watch *model.UserVideoWatchHistory) error {
	if err := DB.WithContext(ctx).Model(&model.UserVideoWatchHistory{}).Create(watch).Error; err != nil {
		return errors.WithMessage(err, "Failed to add UserVideoWatchHistory")
	}
	return nil
}

func AddUserViewBehavior(ctx context.Context, behavior *model.UserBehavior) error {
	if err := DB.WithContext(ctx).Model(&model.UserBehavior{}).Create(behavior).Error; err != nil {
		return errors.WithMessage(err, "Failed to add UserViewBehavior")
	}
	return nil
}

func AddUserShareBehavior(ctx context.Context, behavior *model.UserBehavior) error {
	if err := DB.WithContext(ctx).Model(&model.UserBehavior{}).Create(behavior).Error; err != nil {
		return errors.WithMessage(err, "Failed to add UserShareBehavior")
	}
	return nil
}
