package main

import (
	"context"

	"HuaTug.com/cmd/relation/service"
	"HuaTug.com/kitex_gen/base"
	"HuaTug.com/kitex_gen/relations"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/pkg/errors"
)

type RelationServiceImpl struct{}

func (v *RelationServiceImpl) RelationService(ctx context.Context, req *relations.RelationServiceRequest) (resp *relations.RelationServiceResponse, err error) {
	resp = new(relations.RelationServiceResponse)
	resp.Base = &base.Status{}
	err = service.NewRelationService(ctx).RelationService(ctx, req)
	if err != nil {
		hlog.CtxErrorf(ctx, "service.RelationService failed, original error: %v", errors.Cause(err))
		hlog.CtxErrorf(ctx, "stack trace: \n%+v\n", err)
		resp.Base.Code = consts.StatusBadRequest
		resp.Base.Msg = "Fail to RelationAction!"
		return resp, err
	}
	resp.Base.Code = consts.StatusOK
	resp.Base.Msg = "RelationAction Successfully"
	return resp, nil
}

func (v *RelationServiceImpl) FollowingList(ctx context.Context, req *relations.FollowingListRequest) (resp *relations.FollowingListResponse, err error) {
	resp = new(relations.FollowingListResponse)
	resp, err = service.NewFollowingListService(ctx).FollowingList(ctx, req)
	resp.Base = &base.Status{}
	if err != nil {
		hlog.CtxErrorf(ctx, "service.FollowingList failed, original error: %v", errors.Cause(err))
		hlog.CtxErrorf(ctx, "stack trace: \n%+v\n", err)
		resp.Base.Code = consts.StatusBadRequest
		resp.Base.Msg = "Fail to List Following!"
		return resp, err
	}
	resp.Base.Code = consts.StatusOK
	resp.Base.Msg = "List Following Successfully"
	return resp, nil
}

func (v *RelationServiceImpl) FollowerList(ctx context.Context, req *relations.FollowerListRequest) (resp *relations.FollowerListResponse, err error) {
	resp = new(relations.FollowerListResponse)
	resp, err = service.NewFollowerListService(ctx).FollowerList(ctx, req)
	resp.Base = &base.Status{}
	if err != nil {
		hlog.CtxErrorf(ctx, "service.FollowerList failed, original error: %v", errors.Cause(err))
		hlog.CtxErrorf(ctx, "stack trace: \n%+v\n", err)
		resp.Base.Code = consts.StatusBadRequest
		resp.Base.Msg = "Fail to List Follower!"
		return resp, err
	}
	resp.Base.Code = consts.StatusOK
	resp.Base.Msg = "List Follower Successfully"
	return resp, nil
}

func (v *RelationServiceImpl) FriendList(ctx context.Context, req *relations.FriendListRequest) (resp *relations.FriendListResponse, err error) {
	resp = new(relations.FriendListResponse)
	resp, err = service.NewFriendListService(ctx).FriendList(ctx, req)
	resp.Base = &base.Status{}
	if err != nil {
		hlog.CtxErrorf(ctx, "service.FriendList failed, original error: %v", errors.Cause(err))
		hlog.CtxErrorf(ctx, "stack trace: \n%+v\n", err)
		resp.Base.Code = consts.StatusBadRequest
		resp.Base.Msg = "Fail to List Friend!"
		return resp, err
	}
	resp.Base.Code = consts.StatusOK
	resp.Base.Msg = "List Friend Successfully"
	return resp, nil
}
