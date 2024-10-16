package service

import (
	"context"
	"time"

	"HuaTug.com/cmd/user/dal/db"
	"HuaTug.com/kitex_gen/base"
	"HuaTug.com/kitex_gen/users"
	"HuaTug.com/pkg/constants"
	"HuaTug.com/pkg/utils"
	"github.com/pkg/errors"
)

type CreateUserService struct {
	ctx context.Context
}

func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

func (v *CreateUserService) CreateUser(req *users.CreateUserRequest) error {
	var err error
	var flag bool
	//var wg sync.WaitGroup
	if err, flag = db.RemoveDuplicate(v.ctx, req.UserName); !flag {
		return errors.WithMessage(err, "User duplicate registration")
	}
	passWord, err := utils.Crypt(req.Password)
	if err != nil {
		return errors.WithMessage(err, "Password fail to crypt")
	}
	return db.CreateUser(v.ctx, &base.User{
		Password:  passWord,
		UserName:  req.UserName,
		CreatedAt: time.Now().Format(constants.DataFormate),
		UpdatedAt: time.Now().Format(constants.DataFormate),
		DeletedAt: "",
		AvatarUrl: "HuaTug.com",
	})
}
