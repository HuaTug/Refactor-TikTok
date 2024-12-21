package service

import (
	"context"

	"HuaTug.com/cmd/user/dal/db"
	"HuaTug.com/kitex_gen/users"
)

type CheckUserExistsService struct {
	ctx context.Context
}

func NewCheckUserExistsService(ctx context.Context) *CheckUserExistsService {
	return &CheckUserExistsService{ctx: ctx}
}

func (s *CheckUserExistsService) CheckUserExists(req *users.CheckUserExistsByIdRequst) (exists bool, err error) {
	exists, err = db.CheckUserExistById(s.ctx, req.UserId)
	if err != nil {
		return false, err
	} else if !exists {
		return false, nil
	}
	return true, nil
}
