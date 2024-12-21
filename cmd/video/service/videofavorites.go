package service

import (
	"context"
	"time"

	"HuaTug.com/cmd/model"
	"HuaTug.com/cmd/video/dal/db"
	"HuaTug.com/kitex_gen/base"
	"HuaTug.com/kitex_gen/videos"
	"HuaTug.com/pkg/constants"
	"github.com/pkg/errors"
)

type VideoFavoritesService struct {
	ctx context.Context
}

func NewVideoFavoritesService(ctx context.Context) *VideoFavoritesService {
	return &VideoFavoritesService{
		ctx: ctx,
	}
}

func (s *VideoFavoritesService) CreateFavorite(req *videos.CreateFavoriteRequest) error {
	if err := db.CreateFavorite(s.ctx, &base.Favorite{
		UserId:      req.UserId,
		Name:        req.Name,
		Description: req.Description,
		CoverUrl:    req.CoverUrl,
		CreatedAt:   time.Now().Format(constants.DataFormate),
		DeletedAt:   "",
	}); err != nil {
		return err
	}
	return nil
}

func (s *VideoFavoritesService) GetFavoriteList(req *videos.GetFavoriteListRequest) ([]*base.Favorite, error) {
	var favList []*base.Favorite
	favList, err := db.GetFavoriteList(s.ctx, req)
	if err != nil {
		return favList, errors.WithMessage(err, "Failed to get FavoriteList")
	}
	return favList, nil
}

func (s *VideoFavoritesService) GetFavoriteVideoList(req *videos.GetFavoriteVideoListRequest) ([]*base.Video, error) {
	var video []*base.Video
	video, err := db.GetFavoriteVideoList(s.ctx, req)
	if err != nil {
		return video, errors.WithMessage(err, "Failed to get FavoriteVideoList")
	}
	return video, nil
}

func (s *VideoFavoritesService) GetVideoFromFavorite(req *videos.GetVideoFromFavoriteRequest) (*base.Video, error) {
	var video *base.Video
	video, err := db.GetVideoFromFavorite(s.ctx, req)
	if err != nil {
		return video, errors.WithMessage(err, "Failed to get VideoFromFavorite")
	}
	return video, nil
}

func (s *VideoFavoritesService) AddFavoriteVideo(req *videos.AddFavoriteVideoRequest) error {
	if err := db.AddVideoToFavorite(s.ctx, &model.FavoritesVideos{
		UserId:     req.UserId,
		FavoriteId: req.FavoriteId,
		VideoId:    req.VideoId,
	}); err != nil {
		return errors.WithMessage(err, "Failed to AddFavoriteVideo")
	}
	return nil
}

// 在删除收藏夹的同时 删除视频收藏夹中的视频
func (s *VideoFavoritesService) DeleteFavorite(req *videos.DeleteFavoriteRequest) error {
	if err := db.DeleteFavorite(s.ctx, req); err != nil {
		return errors.WithMessage(err, "Failed to DeleteFavorite")
	}
	return nil
}

func (s *VideoFavoritesService) DeleteVideoFromFavorite(req *videos.DeleteVideoFromFavoriteRequest) error {
	if err := db.DeleteVideoFromFavorite(s.ctx, req); err != nil {
		return errors.WithMessage(err, "Failed to DeleteFavorite")
	}
	return nil
}
