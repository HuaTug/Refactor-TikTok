package db

import (
	"time"

	"HuaTug.com/cmd/model"
	"HuaTug.com/pkg/constants"
)

func CreateFollow(followingId, followersId int64) error {
	if err := DB.Create(&model.Follow{
		FollowingId: followingId,
		FollowersId: followersId,
		CreatedAt:   time.Now().Format(constants.DataFormate),
		DeletedAt:   "",
	}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteFollow(followingId, followersId int64) error {
	if err := DB.Where("following_id = ? And followers_id = ?", followingId, followersId).Delete(&model.Follow{}).Error; err != nil {
		return err
	}
	return nil
}

// following 表示发起关注的用户 follower 表示被关注的用户
// 所以这段代码是在表达获取被用户userId 关注的用户
func GetFollowerListPaged(userId, pageNum, pageSize int64) (*[]int64, error) {
	list := make([]int64, 0)
	if err := DB.Model(&model.Follow{}).Where("following_id = ?", userId).Select("followers_id").Offset(int(pageNum-1) * int(pageSize)).Limit(int(pageSize)).Scan(&list).Error; err != nil {
		return nil, err
	}
	return &list, nil
}

// 这段代码表示所有关注这个用户UserID的用户（粉丝）
func GetFollowingListPaged(userId, pageNum, pageSize int64) (*[]int64, error) {
	list := make([]int64, 0)
	if err := DB.Model(&model.Follow{}).Where("followers_id = ?", userId).Select("following_id").Offset(int(pageNum-1) * int(pageSize)).Limit(int(pageSize)).Scan(&list).Error; err != nil {
		return nil, err
	}
	return &list, nil
}

// 获取朋友的这个业务即为两个用户互相关注对方
func GetFriendListPaged(userId, pageNum, pageSize int64) (*[]int64, error) {
	list := make([]int64, 0)
	if err := DB.Model(&model.Follow{}).Where(`following_id = ? And followers_id in (
	select following_id from follows where followers_id = ?)`, userId, userId).Select("follow_id").Offset(int(pageNum-1) * int(pageSize)).Limit(int(pageSize)).Scan(&list).Error; err != nil {
		return nil, err
	}
	return &list, nil
}

func GetFollowingList(userId int64) (*[]int64, error) {
	list := make([]int64, 0)
	if err := DB.Model(&model.Follow{}).Where("followers_id = ?", userId).Select("following_id").Scan(&list).Error; err != nil {
		return nil, err
	}
	return &list, nil
}

func GetFollowerList(userId int64) (*[]int64, error) {
	list := make([]int64, 0)
	if err := DB.Model(&model.Follow{}).Where("following_id = ?", userId).Select("followers_id").Scan(&list).Error; err != nil {
		return nil, err
	}
	return &list, nil
}

func GetFollowingCount(userId int64) (count int64, err error) {
	if err := DB.Model(&model.Follow{}).Where("following_id = ?", userId).Count(&count).Error; err != nil {
		return -1, err
	}
	return count, nil
}

func GetFollowerCount(userId int64) (count int64, err error) {
	if err := DB.Model(&model.Follow{}).Where("followers_id = ?", userId).Count(&count).Error; err != nil {
		return -1, err
	}
	return count, nil
}

func GetFriendCount(userId int64) (count int64, err error) {
	if err := DB.Model(&model.Follow{}).Where(`following_id = ? And followers_id in (
		select following_id from follows where followers_id = ?
	)`, userId, userId).Count(&count).Error; err != nil {
		return -1, err
	}
	return count, nil
}

func IsRelationExist(followingId, followerId int64) (bool, error) {
	var count int64
	err := DB.Model(&model.Follow{}).Where("following_id = ? And followers_id = ?", followingId, followerId).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
