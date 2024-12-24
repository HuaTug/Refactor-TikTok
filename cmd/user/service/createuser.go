package service

import (
	"context"
	"time"

	"HuaTug.com/cmd/model"
	"HuaTug.com/cmd/user/dal/db"
	"HuaTug.com/kitex_gen/base"
	"HuaTug.com/kitex_gen/users"
	"HuaTug.com/pkg/constants"
	"HuaTug.com/pkg/utils"
	"github.com/cloudwego/hertz/pkg/common/hlog"
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
  hlog.Info("username:",req.UserName,"password:",req.Password,"email:",req.Email)
	err = db.CreateUser(v.ctx, &base.User{
		Password:  passWord,
		UserName:  req.UserName,
		CreatedAt: time.Now().Format(constants.DataFormate),
		UpdatedAt: time.Now().Format(constants.DataFormate),
		Email:     req.Email,
		Sex:       req.Sex,
		AvatarUrl: "HuaTug.com",
	})
	if err != nil {
		return errors.WithMessage(err, "dao.CreateUser failed")
	}

	// 对用户进行角色分配 （实现了对校内外用户的权限划分基础）
	var role_id int64
	var roles string
	if !utils.IsValidEmail(req.Email) {
		roles = "guest"
		role_id, err = db.GetRole(v.ctx, roles)
		if err != nil {
			return errors.WithMessage(err, "dao.GetRole failed")
		}
	} else {
		roles = "user"
		role_id, err = db.GetRole(v.ctx, roles)
		if err != nil {
			return errors.WithMessage(err, "dao.GetRole failed")
		}
	}

	userId, err := db.GetMaxId(v.ctx)
	if err != nil {
		return errors.WithMessage(err, "dao.GetMaxId failed")
	}
	if err = db.Permession_assignment(v.ctx, &model.User_Role{
		Role_Id: role_id,
		User_Id: userId,
		Role:    roles,
	}); err != nil {
		return errors.WithMessage(err, "dao.Permession_assignment failed")
	}
	return nil
}
