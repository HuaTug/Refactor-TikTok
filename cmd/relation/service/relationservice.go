package service

import (
	"context"

	"HuaTug.com/cmd/relation/dal/db"
	"HuaTug.com/kitex_gen/relations"
	"HuaTug.com/pkg/errno"
)

type RelationService struct {
	ctx context.Context
}

func NewRelationService(ctx context.Context) *RelationService {
	return &RelationService{ctx: ctx}
}

func (service *RelationService) RelationService(ctx context.Context, req *relations.RelationServiceRequest) (err error) {
	if req.FromUserId == req.ToUserId {
		return errno.RequestErr
	}
	switch req.ActionType {
	case 1:
		if err := service.CreateFollow(service.ctx, req); err != nil {
			return err
		}
	case 2:
		if err := service.CancleFollow(service.ctx, req); err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (service *RelationService) CreateFollow(ctx context.Context, req *relations.RelationServiceRequest) error {
	if err := db.CreateFollow(req.FromUserId, req.ToUserId); err != nil {
		return errno.ServiceErr
	}
	return nil
}

func (service *RelationService) CancleFollow(ctx context.Context, req *relations.RelationServiceRequest) error {
	if err := db.DeleteFollow(req.FromUserId, req.ToUserId); err != nil {
		return errno.ServiceErr
	}
	return nil
}
