package service

import (
	"context"
	"sync"
	"time"

	"HuaTug.com/cmd/api/rpc"
	"HuaTug.com/cmd/interaction/dal/db"
	"HuaTug.com/cmd/interaction/infras/redis"
	"HuaTug.com/cmd/model"
	"HuaTug.com/kitex_gen/base"
	"HuaTug.com/kitex_gen/interactions"
	"HuaTug.com/kitex_gen/videos"
	"HuaTug.com/pkg/constants"
	"HuaTug.com/pkg/errno"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/pkg/errors"
)

type CommentService struct {
	ctx context.Context
}

func NewCommentService(ctx context.Context) *CommentService {
	return &CommentService{ctx: ctx}
}

func (service *CommentService) CreateComment(ctx context.Context, req *interactions.CreateCommentRequest) (err error) {
	uid := req.UserId
	if req.Content == `` {
		return errno.RequestErr
	}
	if req.CommentId == 0 && req.VideoId == 0 {
		return errno.RequestErr
	}
	// 按照数据表的设计，CommentId不需要传入，此时表示为根评论，但是设置CommentId在传入时表示为想作为某一个评论的子评论
	if req.CommentId == 0 {
		req.CommentId = -1
	} else {
		parentCommentId, err := db.GetParentCommentId(service.ctx, req.CommentId)
		if err != nil {
			return errno.ServiceErr
		}
		// 通过req.Mode完成了 二级评论和三级评论的区分
		if req.Mode != 0 {
			if parentCommentId != 0 {
				//当用户选择评论父评论 且父评论不为空
				req.CommentId = parentCommentId
			}
		}
		// 否则的话 表示为回复评论 此时req.CommentId就表示为子评论的父评论
	}
	if req.VideoId == 0 {
		videoId, err := db.GetCommentVideoId(service.ctx, req.CommentId)
		if err != nil {
			return errors.WithMessage(err, "Failed to get videoId by commentId")
		}
		req.VideoId = videoId
	}

	if err = db.CreateComment(service.ctx, &model.Comment{
		VideoId:   req.VideoId,
		ParentId:  req.CommentId,
		UserId:    uid,
		Content:   req.Content,
		CreatedAt: time.Now().Format(constants.DataFormate),
		UpdatedAt: time.Now().Format(constants.DataFormate),
		DeletedAt: "",
	}); err != nil {
		return errors.WithMessage(err, "Failed to create comment")
	}
	comment := &model.UserBehavior{
		UserId:       uid,
		VideoId:      req.VideoId,
		BehaviorType: "comment",
		BehaviorTime: time.Now().Format(constants.DataFormate),
	}
	go db.AddUserCommentBehavior(service.ctx, comment)
	return nil
}

func (service *CommentService) ListComment(ctx context.Context, req *interactions.ListCommentRequest) (resp *interactions.ListCommentResponse, err error) {
	resp = new(interactions.ListCommentResponse)
	if req.PageNum <= 0 {
		req.PageNum = -1
	}
	if req.PageSize <= 0 {
		req.PageSize = constants.DefaultLimit
	}

	var (
		data *[]*base.Comment
	)
	if req.VideoId != 0 {
		if data, err = service.GetVideoComment(req); err != nil {
			return nil, err
		}
	} else if req.CommentId != 0 {
		if data, err = service.GetCommentComment(req); err != nil {
			return nil, err
		}
	} else {
		return nil, err
	}
	resp.Items = *data
	resp.Base = &base.Status{}
	return resp, nil
}

func (service *CommentService) NewDeleteEvent(ctx context.Context, req *interactions.CommentDeleteRequest) error {
	if req.VideoId != 0 {
		videoInfo, err := rpc.VideoClient.VideoInfo(ctx, &videos.VideoInfoRequest{VideoId: req.VideoId})
		if err != nil {
			hlog.Info("Error in VideoInfo RPC call:", err)
			return errno.RpcErr
		}
		if videoInfo == nil {
			hlog.Error("VideoInfo is nil")
			return errno.ServiceErr
		}
		if videoInfo.Items.UserId != req.FromUserId {
			return errno.ServiceErr
		}
		if err := service.DeleteVideo(req); err != nil {
			return err
		}
	} else if req.CommentId != 0 {
		commentInfo, err := db.GetCommentInfo(service.ctx, req.CommentId)
		if err != nil {
			return errno.MysqlErr
		}
		if commentInfo.UserId != req.FromUserId {
			return errno.ServiceErr
		}
		if err := service.DeleteComment(req); err != nil {
			return err
		}
	} else {
		return errno.RequestErr
	}
	return nil
}

func (service *CommentService) GetVideoComment(req *interactions.ListCommentRequest) (*[]*base.Comment, error) {
	data := make([]*base.Comment, 0)
	list, err := db.GetVideoCommentListByPart(service.ctx, req.VideoId, req.PageNum, req.PageSize)
	if err != nil {
		return nil, errno.ServiceErr
	}
	var (
		wg         sync.WaitGroup
		errChan    = make(chan error, 3)
		res        *model.Comment
		likeCount  int64
		childCount int64
	)
	for _, item := range *list {
		wg.Add(3)
		go func() {
			res, err = db.GetCommentInfo(service.ctx, item)
			if err != nil {
				errChan <- errors.WithMessage(err, "Failed to get CommentInfo")
			}
			wg.Done()
		}()
		go func() {
			likeCount, err = redis.GetVideoLikeCount(item)
			if err != nil {
				errChan <- errors.WithMessage(err, "Failed to get VideoVisitCount")
			}
			wg.Done()
		}()
		go func() {
			childCount, err = db.GetChildCommentCount(service.ctx, item)
			if err != nil {
				errChan <- errors.WithMessage(err, "Failed to get ChildCommentCount")
			}
			wg.Done()
		}()
		wg.Wait()
		select {
		case result := <-errChan:
			return nil, result
		default:
		}
		data = append(data, &base.Comment{
			CommentId:  res.CommentId,
			VideoId:    res.VideoId,
			UserId:     res.UserId,
			ParentId:   res.ParentId,
			LikeCount:  likeCount,
			ChildCount: childCount,
			Content:    res.Content,
			CreatedAt:  res.CreatedAt,
			UpdatedAt:  res.UpdatedAt,
			DeletedAt:  res.DeletedAt,
		})
	}
	return &data, nil
}

func (service *CommentService) GetCommentComment(req *interactions.ListCommentRequest) (*[]*base.Comment, error) {
	data := make([]*base.Comment, 0)
	list, err := db.GetCommentChildListByPart(service.ctx, req.CommentId, req.PageNum, req.PageSize)
	if err != nil {
		return nil, errno.ServiceErr
	}
	var (
		wg         sync.WaitGroup
		errChan    = make(chan error, 3)
		res        *model.Comment
		likeCount  int64
		childCount int64
	)
	for _, item := range *list {
		wg.Add(3)
		go func() {
			res, err = db.GetCommentInfo(service.ctx, item)

			if err != nil {
				errChan <- errors.WithMessage(err, "Failed to get CommentInfo")
			}
			wg.Done()
		}()
		go func() {
			likeCount, err = redis.GetCommentLikeCount(item)
			if err != nil {
				errChan <- errors.WithMessage(err, "Failed to get VideoVisitCount")
			}
			wg.Done()
		}()
		go func() {
			childCount, err = db.GetChildCommentCount(service.ctx, item)
			if err != nil {
				errChan <- errors.WithMessage(err, "Failed to get ChildCommentCount")
			}
			wg.Done()
		}()
		wg.Wait()
		select {
		case result := <-errChan:
			return nil, result
		default:
		}
		data = append(data, &base.Comment{
			CommentId:  res.CommentId,
			VideoId:    res.VideoId,
			UserId:     res.UserId,
			ParentId:   res.ParentId,
			LikeCount:  likeCount,
			ChildCount: childCount,
			Content:    res.Content,
			CreatedAt:  res.CreatedAt,
			UpdatedAt:  res.UpdatedAt,
			DeletedAt:  res.DeletedAt,
		})
	}
	return &data, nil
}

func (service *CommentService) DeleteVideo(req *interactions.CommentDeleteRequest) error {
	list, err := db.GetVideoCommentList(context.Background(), req.VideoId)
	if err != nil {
		return errno.MysqlErr
	}
	if _, err := rpc.VideoClient.VideoDelete(service.ctx, &videos.VideoDeleteRequest{VideoId: req.VideoId, UserId: req.FromUserId}); err != nil {
		return errno.ServiceErr
	}

	var (
		wg      sync.WaitGroup
		errChan = make(chan error, len(*list))
	)

	wg.Add(len(*list))
	for _, item := range *list {
		go func(commentId int64) {
			if err := service.DeleteComment(&interactions.CommentDeleteRequest{CommentId: commentId}); err != nil {
				errChan <- err
			}
			wg.Done()
		}(item)
	}

	wg.Wait()
	select {
	case result := <-errChan:
		return result
	default:
	}
	return nil

}

func (service *CommentService) DeleteComment(req *interactions.CommentDeleteRequest) error {
	if err := db.DeleteComment(service.ctx, req.CommentId); err != nil {
		return errno.ServiceErr
	}

	var (
		wg      sync.WaitGroup
		errChan = make(chan error, 2)
	)
	wg.Add(2)
	go func() {
		if err := db.DeleteComment(context.Background(), req.CommentId); err != nil {
			errChan <- errno.RedisErr
		}
		wg.Done()
	}()
	go func() {
		if err := redis.DeleteCommentAndAllAbout(req.CommentId); err != nil {
			errChan <- errno.RedisErr
		}
		wg.Done()
	}()
	wg.Wait()
	select {
	case errr := <-errChan:
		return errr
	default:
	}
	return nil
}

func (service *CommentService) NewVideoPopularListEvent(req *interactions.VideoPopularListRequest) (*[]string, error) {
	list, err := redis.GetVideoPopularList(req.PageNum, req.PageSize)
	if err != nil {
		return nil, errno.RedisErr
	}
	return list, nil
}

// 删除操作
func (service *CommentService) NewDeleteVideoInfoEvent(req *interactions.DeleteVideoInfoRequest) error {
	var (
		err         error
		commentList *[]int64
		wg          sync.WaitGroup
		errChan     = make(chan error, 1)
	)
	// 将查询操作放入外面 保证数据的一致性
	if commentList, err = db.GetVideoCommentList(service.ctx, req.VideoId); err != nil {
		return errors.New("Failed to get VideoCommentList")
	}

	wg.Add(2)
	// 删除评论
	go func() {
		defer wg.Done()
		if err := redis.DeleteAllComment(*commentList); err != nil {
			errChan <- errors.New("Failed to delete VideoComment")
			return
		}
	}()
	// 删除点赞
	go func() {
		defer wg.Done()
		if err := redis.DeleteVideoAndAllAbout(req.VideoId); err != nil {
			errChan <- errors.New("Failed to delete VideoLike")
			return
		}
	}()
	wg.Wait()

	select {
	case err = <-errChan:
		return err
	default:
		return nil
	}
}
