package common

import (
	"context"
	"strconv"
	"time"

	"HuaTug.com/cmd/interaction/dal/db"
	"HuaTug.com/cmd/interaction/infras/redis"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type CommentSync struct {
	ctx    context.Context
	cancle context.CancelFunc
}

type CommentSyncData struct {
	commentId int64
	likelist  *[]int64
}

func NewCommentSync() *CommentSync {
	ctx, cancle := context.WithCancel(context.Background())
	return &CommentSync{
		ctx:    ctx,
		cancle: cancle,
	}
}
func (sm *CommentSync) Run() {
	hlog.Info("comment_run start")
	if err := commentSyncInit(); err != nil {
		hlog.Info(err)
		return
	}
	go func() {
		for {
			time.Sleep(time.Second * 10)
			select {
			case <-sm.ctx.Done():
				hlog.Info("Ok,stop CommentSync!")
				return
			default:
			}

			commentIdList, err := db.GetCommentIdList(context.Background())
			if err != nil {
				hlog.Warn(err)
			}

			for _, cid := range *commentIdList {
				likelist, err := redis.GetNewUpdateCommentLikeList(cid)
				if err != nil {
					hlog.Info(err)
					continue
				}

				for _, uid := range *likelist {
					userID, _ := strconv.ParseInt(uid, 10, 64)
					if err := db.CreateCommentLike(context.Background(), cid, userID); err != nil {
						hlog.Error(err)
					}
					
					if err := redis.AppendCommentLikeInfoToStaticSpace(cid, userID); err != nil {
						hlog.Error(err)
					}
					// 删除动态点赞
					if err := redis.DeleteCommentLikeInfoFromDynamicSpace(cid, userID); err != nil {
						hlog.Error(err)
					}
				}

				dislikeList, err := redis.GetNewDeleteCommentLikeList(cid)
				if err != nil {
					hlog.Error(err)
					continue
				}
				for _, uid := range *dislikeList {
					userId, _ := strconv.ParseInt(uid, 10, 64)
					if err := db.DeleteCommentLike(context.Background(), cid, userId); err != nil {
						hlog.Info(err)
					}
					if err := redis.DeleteCommentLikeInfoFromDynamicSpace(cid, userId); err != nil {
						hlog.Error(err)
					}
				}
			}
		}
	}()
}

func (sm *CommentSync) Stop() {
	sm.cancle()
}

func commentSyncInit() error {
	list, err := db.GetCommentIdList(context.Background())
	if err != nil {
		hlog.Info(err)
		panic(err)
	}

	var (
		syncList = make([]CommentSyncData, 0)
		data     CommentSyncData
	)
	for _, v := range *list {
		data.commentId = v
		if data.likelist, err = db.GetCommentLikeList(context.Background(), v); err != nil {
			hlog.Info(err)
			return err
		}
		syncList = append(syncList, data)
	}
	if err := CommentSyncDB2Redis(&syncList); err != nil {
		hlog.Info(err)
		return err
	}
	return nil
}

func CommentSyncDB2Redis(syncList *[]CommentSyncData) error {
	for _, item := range *syncList {
		if err := redis.PutCommentLikeInfo(item.commentId, item.likelist); err != nil {
			return err
		}
	}
	return nil
}
