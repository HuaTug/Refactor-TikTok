package redis

import (
	"context"
	"fmt"
	"sync"

	"HuaTug.com/cmd/interaction/dal/db"
	"github.com/go-redis/redis"
)

// 保证数据的一致性 (同步之前先把原来的数据删除)
func PutVideoLikeInfo(videoId int64, uidList *[]string) error {
	pipe := redisDBVideoInfo.TxPipeline()
	pipe.Del("l_video:" + fmt.Sprint(videoId))
	pipe.Del("nl_video:" + fmt.Sprint(videoId))
	for _, item := range *uidList {
		pipe.SAdd("l_video:"+fmt.Sprint(videoId), item)
	}
	if _, err := pipe.Exec(); err != nil {
		return err
	}
	return nil
}

// 表示为动态点赞（只有点或不点，即为1或2这两种情况）
func AppendVideoLikeInfo(videoId, userId int64) error {
	if _, err := redisDBVideoInfo.ZAdd("nl_video:"+fmt.Sprint(videoId), redis.Z{Score: 1, Member: userId}).Result(); err != nil {
		return err
	}

	// 避免发生重复点赞的情况出现
	if _, err := redisDBVideoInfo.SRem("l_video:"+fmt.Sprint(videoId), userId).Result(); err != nil {
		return err
	}
	return nil
}

// 这段代码是用来将用户点赞的视频信息添加到静态空间中，使用SET集合完成存储用户点赞的视频以及该用户的信息（以实现存储谁点赞了哪条视频）
func AppendVideoLikeInfoToStaticSpace(videoId, userId int64) error {
	if _, err := redisDBVideoInfo.SAdd("l_video:"+fmt.Sprint(videoId), userId).Result(); err != nil {
		return err
	}
	return nil
}

func DeleteVideoLikeInfoFromDynamicSpace(videoId, userId int64) error {
	if _, err := redisDBVideoInfo.ZRem("nl_video:"+fmt.Sprint(videoId), userId).Result(); err != nil {
		return err
	}
	return nil
}

func RemoveVideoLikeInfo(videoId, userId int64) error {
	if _, err := redisDBVideoInfo.ZAdd("nl_video:"+fmt.Sprint(videoId), redis.Z{Score: 2, Member: userId}).Result(); err != nil {
		return err
	}
	if _, err := redisDBVideoInfo.SRem("l_video:"+fmt.Sprint(videoId), fmt.Sprint(userId)).Result(); err != nil {
		return err
	}
	return nil
}

// 这段代码的作用是用来检查某个用户是否已经对特定的视频点赞
func IsVideoLikedByUser(videoId, userId int64) (bool, error) {
	//redis集合的判断函数
	exists, err := redisDBVideoInfo.SIsMember("l_video:"+fmt.Sprint(videoId), userId).Result()
	if err != nil {
		return false, err
	}
	if !exists {
		//在Redis中可以使用ZScore或者是ZRank获得在ZSet中某一个值得分数或者是排名，以此来判断这个值是否存在在ZSet集合中
		_, err := redisDBVideoInfo.ZScore("nl_video:"+fmt.Sprint(videoId), fmt.Sprint(userId)).Result()
		if err != nil {
			return false, err
		}
		return true, nil
	} else {
		return true, nil
	}
}

// 这段代码是将两种点赞状态进行了求和
func GetVideoLikeList(videoId int64) (*[]string, error) {
	//SMebers用于获取集合Key的所有元素
	list, err := redisDBVideoInfo.SMembers("l_video:" + fmt.Sprint(videoId)).Result()
	if err != nil {
		return nil, err
	}
	nlist, err := GetNewUpdateVideoLikeList(videoId)
	if err != nil {
		return nil, err
	}
	list = append(list, *nlist...)
	return &list, nil
}

// 下面的两端代码是用来获取动态点赞和取消点赞的用户列表（用1和2进行区分）
func GetNewUpdateVideoLikeList(videoId int64) (*[]string, error) {
	//这个函数可以获得进行动态点赞的用户
	list, err := redisDBVideoInfo.ZRangeByScore("nl_video:"+fmt.Sprint(videoId), redis.ZRangeBy{Min: "1", Max: "1"}).Result()
	if err != nil {
		return nil, err
	}
	return &list, nil
}

func GetNewDeleteVideoLikeList(videoId int64) (*[]string, error) {
	list, err := redisDBVideoInfo.ZRangeByScore("nl_video:"+fmt.Sprint(videoId), redis.ZRangeBy{Min: "2", Max: "2"}).Result()
	if err != nil {
		return nil, err
	}
	return &list, nil
}

// 用于获得动态点赞和静态点赞的总数
func GetVideoLikeCount(videoId int64) (int64, error) {
	var count int64
	var ncount int64
	var err error
	if count, err = redisDBVideoInfo.SCard("l_video:" + fmt.Sprint(videoId)).Result(); err != nil {
		return -1, err
	}
	if ncount, err = redisDBVideoInfo.ZCount("nl_video:"+fmt.Sprint(videoId), "1", "1").Result(); err != nil {
		return -1, err
	}
	return count + ncount, nil
}

// 用于获得视频的热度排行（根据pageNum和pageSize进行调节）
func GetVideoPopularList(pageNum, pageSize int64) (*[]string, error) {
	list, err := redisDBVideoInfo.ZRevRange("visit", (pageNum-1)*pageSize, pageNum*pageSize).Result()
	if err != nil {
		return nil, err
	}
	return &list, nil
}

func DeleteVideoAndAllAbout(videoId int64) error {
	videoPipe := redisDBVideoInfo.TxPipeline()
	commentPipe := redisDBCommentInfo.TxPipeline()
	commenList, err := db.GetVideoCommentList(context.Background(), videoId)
	if err != nil {
		return err
	}

	videoPipe.Del("nl_video:" + fmt.Sprint(videoId))
	videoPipe.Del("l_video:" + fmt.Sprint(videoId))

	for _, item := range *commenList {
		commentPipe.Del("l_video:" + fmt.Sprint(item))
		commentPipe.Del("nl_video:" + fmt.Sprint(item))
	}

	var (
		wg      sync.WaitGroup
		errChan = make(chan error, 2)
	)
	wg.Add(2)
	go func() {
		if _, err := videoPipe.Exec(); err != nil {
			errChan <- err
		}
		wg.Done()
	}()

	go func() {
		if _, err := commentPipe.Exec(); err != nil {
			errChan <- err
		}
		wg.Done()
	}()
	wg.Wait()
	select {
	case result := <-errChan:
		return result
	default:
	}
	return nil
}
