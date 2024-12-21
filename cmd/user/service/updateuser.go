package service

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"HuaTug.com/cmd/user/dal/db"
	"HuaTug.com/kitex_gen/base"
	"HuaTug.com/kitex_gen/users"
	"HuaTug.com/pkg/constants"
	"HuaTug.com/pkg/errno"
	"HuaTug.com/pkg/oss"
	"github.com/pkg/errors"
)

type UpdateUserService struct {
	ctx context.Context
}

func NewUpdateUserService(ctx context.Context) *UpdateUserService {
	return &UpdateUserService{ctx: ctx}
}

func (v *UpdateUserService) UpdateUser(req *users.UpdateUserRequest) (err error) {
	if avatarUrl, err := v.uploadAvatarToOss(fmt.Sprint(req.UserId), req.Data, req.Filesize); err != nil {
		return errors.WithMessage(err, "DataProcess failed")
	} else {
		user := &base.User{
			UserId:    req.UserId,
			UserName:  req.UserName,
			Password:  req.Password,
			UpdatedAt: time.Now().Format(constants.DataFormate),
			AvatarUrl: avatarUrl,
		}
		if err := db.UpdateUser(v.ctx, user); err != nil {
			return errors.WithMessage(err, "dao.UpdateUser failed")
		}
	}
	return nil
}

func (v *UpdateUserService) uploadAvatarToOss(uid string, fileContent []byte, fileSize int64) (string, error) {
	fileType := http.DetectContentType(fileContent)
	switch fileType {
	case `image/png`, `image/jpg`, `image/jpeg`:
		{
			var avatarUrl string
			var err error
			if avatarUrl, err = oss.UploadAvatar(&fileContent, fileSize, uid, fileType); err != nil {
				return "", errno.ServiceErr
			}
			return avatarUrl, nil
		}
	default:
		return "", errno.DataProcessErr
	}
}
