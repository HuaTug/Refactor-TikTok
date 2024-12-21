package db

import (
	"context"

	"HuaTug.com/cmd/model"
)

func Permession_assignment(ctx context.Context, permession *model.User_Role) error {
	err := DB.WithContext(ctx).Model(&model.User_Role{}).Create(permession).Error
	if err != nil {
		return err
	}
	return nil
}

func GetRole(ctx context.Context, role string) (int64, error) {
	var roleId int64
	err := DB.WithContext(ctx).Model(&model.Role{}).Where("role = ?", role).Select("role_id").Scan(&roleId).Error
	if err != nil {
		return 0, err
	}
	return roleId, nil
}
