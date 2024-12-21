package db

import (
	"context"
	"time"

	"HuaTug.com/cmd/model"
	"HuaTug.com/pkg/constants"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/google/uuid"
)

func CreateComment(ctx context.Context, comment *model.Comment) error {
	return DB.WithContext(ctx).Create(comment).Error
}

func DeleteComment(ctx context.Context, commentId int64) error {
	if err := DB.WithContext(ctx).Model(&model.Comment{}).Where("comment_id = ?", commentId).Delete(&model.Comment{}).Error; err != nil {
		return err
	}
	return nil
}

// 获取子评论的数目
func GetChildCommentCount(ctx context.Context, commentId int64) (int64, error) {
	var count int64
	if err := DB.WithContext(ctx).Model(&model.Comment{}).Where("parent_id = ?", commentId).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func GetVideoCommentCount(ctx context.Context, videoId int64) (count int64, err error) {
	if err := DB.WithContext(ctx).Model(&model.Comment{}).Where("video_id = ?", videoId).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// 获取某一条评论的全部信息
func GetCommentInfo(ctx context.Context, commentId int64) (comment *model.Comment, err error) {
	comment = &model.Comment{}
	if err := DB.WithContext(ctx).Model(&model.Comment{}).Where("comment_id = ?", commentId).Find(comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

func GetParentCommentId(ctx context.Context, commentId int64) (parentId int64, err error) {
	if err := DB.WithContext(ctx).Model(&model.Comment{}).Where("comment_id = ?", commentId).Select("parent_id").Find(&parentId).Error; err != nil {
		return 0, err
	}
	return parentId, nil
}

// 获得点赞某一条评论的所有用户
func GetCommentLikeList(ctx context.Context, commentId int64) (*[]int64, error) {
	list := make([]int64, 0)
	//ToDo :sql语句的执行顺序对于性能有无影响
	if err := DB.WithContext(ctx).Model(&model.CommentLike{}).Where("comment_id = ?", commentId).Select("user_id").Scan(&list).Error; err != nil {
		return nil, err
	}
	return &list, nil
}

// 这段代码表示获得评论的点赞数
func GetCommentLikeCount(ctx context.Context, commentId int64) (count int64, err error) {
	if err := DB.WithContext(ctx).Model(&model.CommentLike{}).Where("comment_id = ?", commentId).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// 获得被评论的视频Id
func GetCommentVideoId(ctx context.Context, commentId int64) (videoId int64, err error) {
	if err := DB.WithContext(ctx).Model(&model.Comment{}).Where("comment_id = ?", commentId).Select("video_id").Find(&videoId).Error; err != nil {
		return 0, err
	}
	return videoId, nil
}

// 获取子评论列表
func GetCommentChildList(ctx context.Context, comment_id int64) (*[]int64, error) {
	list := make([]int64, 0)
	if err := DB.WithContext(ctx).Model(&model.Comment{}).Where("parent_id = ?", comment_id).Select("comment_id").Scan(&list).Error; err != nil {
		return nil, err
	}
	return &list, nil
}

func GetCommentChildListByPart(ctx context.Context, comment_id, pagenum, pagesize int64) (*[]int64, error) {
	list := make([]int64, 0)
	if err := DB.WithContext(ctx).Model(&model.Comment{}).Where("parent_id = ?", comment_id).Select("comment_id").Scan(&list).Limit(int(pagesize)).Offset(int(pagenum-1) * int(pagesize)).Error; err != nil {
		return nil, err
	}
	return &list, nil
}

// 获取视频的评论列表
func GetVideoCommentList(ctx context.Context, video_id int64) (*[]int64, error) {
	list := make([]int64, 0)
	if err := DB.WithContext(ctx).Model(&model.Comment{}).Where("video_id = ?", video_id).Select("comment_id").Scan(&list).Error; err != nil {
		return nil, err
	}
	return &list, nil
}

func GetVideoCommentListByPart(ctx context.Context, video_id, pagenum, pagesize int64) (*[]int64, error) {
	list := make([]int64, 0)
	if err := DB.WithContext(ctx).Model(&model.Comment{}).Where("video_id = ?", video_id).Select("comment_id").Scan(&list).Limit(int(pagesize)).Offset(int(pagenum-1) * int(pagesize)).Error; err != nil {
		return nil, err
	}
	return &list, nil
}

func GetCommentIdList(ctx context.Context) (*[]int64, error) {
	list := make([]int64, 0)
	if err := DB.WithContext(ctx).Model(&model.Comment{}).Select("comment_id").Scan(&list).Error; err != nil {
		return nil, err
	}
	return &list, nil
}

// 用来检查给定的comment_id是否在这个数据表中
func IsCommentIdList(ctx context.Context, comment_id int64) (bool, error) {
	var count int64
	if err := DB.WithContext(ctx).Model(&model.Comment{}).Where("comment_id = ?", comment_id).Count(&count).Error; err != nil {
		return false, err
	}
	return count != 0, nil
}

func CreateCommentLike(ctx context.Context, comemntId, userId int64) error {
	uuid := uuid.New().ID()
	if err := DB.WithContext(ctx).Create(&model.CommentLike{
		CommentLikesId: int64(uuid),
		CommentId:      comemntId,
		UserId:         userId,
		CreatedAt:      time.Now().Format(constants.DataFormate),
		DeletedAt:      "",
	}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteCommentLike(ctx context.Context, commentId, UserId int64) error {
	if err := DB.WithContext(ctx).Model(&model.CommentLike{}).Where("comment_id = ? And user_id = ?", commentId, UserId).Delete(&model.CommentLike{}).Error; err != nil {
		hlog.Info(err)
		return err
	}
	return nil
}

func CreateVideoLike(ctx context.Context, videoLike *model.VideoLike) error {
	if err := DB.Create(videoLike).Error; err != nil {
		return err
	}
	return nil
}

func DeleteVideoLike(ctx context.Context, videoId, userId int64) error {
	if err := DB.WithContext(ctx).Model(&model.VideoLike{}).Where("video_id = ? And user_id = ?", videoId, userId).Delete(&model.VideoLike{}).Error; err != nil {
		return err
	}
	return nil
}

// 与下面的函数刚好相对应，即这个函数获得是这个视频被多少人喜欢，下面的函数则是一个用户喜欢的所有视频
func GetVideoLikeList(ctx context.Context, videoId int64) (*[]string, error) {
	list := make([]string, 0)
	if err := DB.WithContext(ctx).Model(&model.VideoLike{}).Where("video_id = ?", videoId).Select("user_id").Scan(&list).Error; err != nil {
		return nil, err
	}
	return &list, nil
}

// 获取用户喜欢的视频列表
func GetVideoLikeListByUserId(ctx context.Context, userId, pageNum, pageSize int64) (*[]string, error) {
	list := make([]string, 0)
	if err := DB.WithContext(ctx).Model(&model.VideoLike{}).Where("user_id = ?", userId).Select("video_id").Offset(int(pageNum-1) * int(pageSize)).Limit(int(pageSize)).Select("video_id").Scan(&list).Error; err != nil {
		return nil, err
	}
	return &list, nil
}

// 记录用户的点赞行为
func AddUserLikeBehavior(ctx context.Context, behavior *model.UserBehavior) error {
	if err := DB.WithContext(ctx).Create(behavior).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUserLikeBehavior(ctx context.Context, userId, videoId int64, behavior string) error {
	if err := DB.WithContext(ctx).Model(&model.UserBehavior{}).Where("user_id = ? and video_id = ? and behavior_type = ?", userId, videoId, behavior).Delete(&model.UserBehavior{}).Error; err != nil {
		return err
	}
	return nil
}

// 记录用户的评论行为
func AddUserCommentBehavior(ctx context.Context, behavior *model.UserBehavior) error {
	if err := DB.WithContext(ctx).Create(behavior).Error; err != nil {
		return err
	}
	return nil
}
