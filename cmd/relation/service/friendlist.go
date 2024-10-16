package service

import (
	"context"

	"HuaTug.com/cmd/relation/dal/db"
	"HuaTug.com/cmd/relation/infras"
	"HuaTug.com/kitex_gen/base"
	"HuaTug.com/kitex_gen/relations"
	"HuaTug.com/kitex_gen/users"
	"HuaTug.com/pkg/constants"
	"HuaTug.com/pkg/errno"
)

type FriendListService struct {
	ctx context.Context
}

func NewFriendListService(ctx context.Context) *FriendListService {
	return &FriendListService{ctx: ctx}
}

func (service *FriendListService) FriendList(ctx context.Context, req *relations.FriendListRequest) (resp *relations.FriendListResponse, err error) {
	resp = new(relations.FriendListResponse)
	userInfo, err := infras.UserClient.GetUserInfo(service.ctx, &users.GetUserInfoRequest{UserId: req.UserId})
	if err != nil {
		return nil, err
	}

	if userInfo == nil {
		return nil, errno.UserNotExistErr
	}
	if req.PageNum <= 0 {
		req.PageNum = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = constants.DefaultLimit
	}

	list, err := db.GetFriendListPaged(req.UserId, req.PageNum, req.PageSize)
	if err != nil {
		return nil, errno.ServiceErr
	}
	data := make([]*base.UserLite, 0)
	for _, item := range *list {
		userInfo, err := infras.UserClient.GetUserInfo(service.ctx, &users.GetUserInfoRequest{UserId: item})
		if err != nil {
			return nil, errno.ServiceErr
		}

		d := &base.UserLite{
			Uid:       item,
			UserName:  userInfo.User.UserName,
			AvatarUrl: userInfo.User.AvatarUrl,
		}
		data = append(data, d)
	}
	total, err := db.GetFriendCount(req.UserId)
	if err != nil {
		return nil, errno.ServiceErr
	}
	resp.Items = data
	resp.Total = total
	return resp, nil
}
