package redis

import (
	"context"
	"fmt"

	"HuaTug.com/cmd/interaction/dal/db"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

func PutCommentLikeInfo(commentId int64, userIdList *[]int64) error {
	//创建管道
	pipe := redisDBCommentInfo.TxPipeline()
	pipe.Del("l_comment:" + fmt.Sprint(commentId))
	pipe.Del("nl_comment:" + fmt.Sprint(commentId))
	for _, item := range *userIdList {
		pipe.SAdd("l_comment:"+fmt.Sprint(commentId), item)
	}
	if _, err := pipe.Exec(); err != nil {
		return errors.WithMessage(err, "Pipe Exec failed")
	}
	return nil
}

func GetCommentLikeCount(commentId int64) (int64, error) {
	//	SCard用于获取集合中的元素数量
	countOld, err := redisDBCommentInfo.SCard("l_comment:" + fmt.Sprint(commentId)).Result()
	if err != nil {
		return -1, err
	}
	//	
	countNew, err := redisDBCommentInfo.ZCount("nl_comment:"+fmt.Sprint(commentId), "1", "1").Result()
	if err != nil {
		return -1, err
	}
	return countOld + countNew, nil
}

func GetCommentLikeList(commentId int64) (*[]string, error) {
	list, err := redisDBCommentInfo.SMembers("l_comment:" + fmt.Sprint(commentId)).Result()
	if err != nil {
		return nil, err
	}
	nl_commentist, err := GetNewUpdateCommentLikeList(commentId)
	if err != nil {
		return nil, err
	}
	list = append(list, *nl_commentist...)
	return &list, nil
}

func GetNewUpdateCommentLikeList(commentId int64) (*[]string, error) {
	list, err := redisDBCommentInfo.ZRangeByScore("nl_comment:"+fmt.Sprint(commentId), redis.ZRangeBy{Min: "1", Max: "1"}).Result()
	if err != nil {
		return nil, err
	}
	return &list, err
}

func GetNewDeleteCommentLikeList(commentId int64) (*[]string, error) {
	//这个过程类似于点赞过程发生一次 然后取消点赞过程发生一次 所有是两次
	list, err := redisDBCommentInfo.ZRangeByScore("nl_comment:"+fmt.Sprint(commentId), redis.ZRangeBy{Min: "2", Max: "2"}).Result()
	if err != nil {
		return nil, err
	}
	return &list, err
}

func AppendCommentLikeInfo(commentId, userId int64) error {
	exists, err := redisDBCommentInfo.ZCount("nl_comment:"+fmt.Sprint(commentId), "1", "1").Result()
	if err != nil && err != redis.Nil {
		return err
	}
	if exists != 0 {
		return fmt.Errorf("user:%d has already liked video:%d", userId, commentId)
	}
	if _, err := redisDBCommentInfo.ZAdd("nl_comment:"+fmt.Sprint(commentId), redis.Z{Score: 1, Member: userId}).Result(); err != nil {
		return err
	}
	_, err = redisDBCommentInfo.ZAdd("nl_comment:"+fmt.Sprint(commentId), redis.Z{Score: 1, Member: userId}).Result()
	if err != nil {
		return err
	}
	if _, err := redisDBCommentInfo.SRem("l_comment:"+fmt.Sprint(commentId), userId).Result(); err != nil {
		return err
	}
	return nil
}

func AppendCommentLikeInfoToStaticSpace(commentId, userId int64) error {
	if _, err := redisDBCommentInfo.SAdd("l_comment:"+fmt.Sprint(commentId), userId).Result(); err != nil {
		return err
	}
	return nil
}

// ToDo
func DeleteCommentLikeInfoFromDynamicSpace(commentId, userId int64) error {
	if _, err := redisDBCommentInfo.ZRem("nl_comment:"+fmt.Sprint(commentId), userId).Result(); err != nil {
		return err
	}
	return nil
}

func RemoveCommentLikeInfo(commentId, userId int64) error {
	exists, err := redisDBCommentInfo.ZCount("nl_comment:"+fmt.Sprint(commentId), "2", "2").Result()
	if err != nil && err != redis.Nil {
		return err
	}
	if exists != 0 {
		return fmt.Errorf("user:%d has already liked video:%d", userId, commentId)
	}
	if _, err := redisDBCommentInfo.ZAdd("nl_comment:"+fmt.Sprint(commentId), redis.Z{Score: 1, Member: userId}).Result(); err != nil {
		return err
	}
	_, err = redisDBCommentInfo.ZAdd("nl_comment:"+fmt.Sprint(commentId), redis.Z{Score: 2, Member: userId}).Result()
	if err != nil {
		return err
	}
	if _, err := redisDBCommentInfo.SRem("l_comment:"+fmt.Sprint(commentId), userId).Result(); err != nil {
		return err
	}
	return nil
}

// 删除所有评论
func DeleteAllComment(commentIdList []int64) error {
	var (
		childList   *[]int64
		err         error
		commentPipe = redisDBCommentInfo.TxPipeline()
	)

	for _, commentId := range commentIdList {
		commentPipe.Unlink("l_comment:" + fmt.Sprint(commentId))
		commentPipe.Unlink("nl_comment:" + fmt.Sprint(commentId))

		if childList, err = db.GetCommentChildList(context.Background(), commentId); err != nil {
			return err
		}

		for _, item := range *childList {
			commentPipe.Unlink("l_comment:" + fmt.Sprint(item))
			commentPipe.Unlink("nl_comment:" + fmt.Sprint(item))
		}
	}

	if _, err := commentPipe.Exec(); err != nil {
		return err
	}

	return nil
}

// 删除一条评论及其子评论
func DeleteCommentAndAllAbout(commentId int64) error {
	return DeleteAllComment([]int64{commentId})
}
