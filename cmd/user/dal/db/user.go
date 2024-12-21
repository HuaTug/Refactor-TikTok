package db

import (
	"context"

	"HuaTug.com/cmd/model"
	"HuaTug.com/kitex_gen/base"
	"HuaTug.com/pkg/utils"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const TableName = "base"

/* func (u *User) TableName() string {
	return TableName
} */

/*
User 结构体有一个 TableName 方法，返回的是 "base"。所以当你在 CreateUser 函数中调用 DB.WithContext(ctx).Create(base).Error 时，
GORM 会自动调用 base（即 *User 类型）的 TableName 方法，得到 "base"，然后在 "base" 这个表上执行插入操作。
*/

func CreateUser(ctx context.Context, user *base.User) error {
	if err := DB.WithContext(ctx).Model(&base.User{}).Create(user); err != nil {
		//ToDo:分层处理err
		return errors.Wrapf(err.Error, "Creareq.UserIdteUser failed,err: %v", err)
	}
	return nil
}

func CheckUser(ctx context.Context, username, password string) (base.User, error, bool) {

	var user base.User
	var count int64
	if err := DB.WithContext(ctx).Model(&base.User{}).Where("Binary user_name=?", username).Count(&count).Find(&user).Error; err != nil {
		hlog.Info(err)
		return user, errors.Wrap(err, "查询用户存在性失败"), false
	}
	if count == 0 {
		logrus.Info("没有这个用户，请重新登录")
		return user, errors.Errorf("数据库中不存在这个用户"), false
	} else if count == 1 {
		if err, flag := utils.VerifyPassword(password, user.Password); !flag {
			return user, errors.Wrapf(err, "Password Wrong,err:%v", err), false
		}
	}
	return user, nil, true
}

func CheckUserExistById(ctx context.Context, userId int64) (bool, error) {
	var user base.User
	if err := DB.WithContext(ctx).Where("user_id=?", userId).Find(&user).Error; err != nil {
		return false, errors.Wrapf(err, "User not exist,err:%v", err)
	}
	if user == (base.User{}) {
		return false, nil
	}
	return true, nil
}

func DeleteUser(ctx context.Context, userId int64) error {
	if err := DB.WithContext(ctx).Where("user_id = ?", userId).Delete(&base.User{}).Error; err != nil {
		return errors.Wrapf(err, "Delete user failed, userId: %d", userId)
	}
	return nil
}

func UpdateUser(ctx context.Context, user *base.User) error {
	if err := DB.WithContext(ctx).Model(&base.User{}).Where("user_id=?", user.UserId).Updates(user).Error; err != nil {
		return errors.Wrapf(err, "Update user failed,err: %v", err)
	}
	return nil
}

func QueryUser(ctx context.Context, keyword *string, page, pageSize int64) ([]*base.User, int64, error) {
	db := DB.WithContext(ctx).Model(base.User{})
	if keyword != nil && len(*keyword) != 0 {
		db = db.Where("Binary user_name like ?", "%"+*keyword+"%")
	}
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, total, errors.Wrapf(err, "QueryUser count failed,err:%v", err)
	}
	var res []*base.User
	if err := db.Limit(int(pageSize)).Offset(int(pageSize * (page - 1))).Find(&res); err != nil || total == 0 {
		return res, total, errors.Wrapf(err.Error, "Limit failed,err:%v", err)
	}
	// for _,u :=range temp{
	// 	res=append(res, &model.User{
	// 		UserId: u.UserId,
	// 		UserName: u.UserName,
	// 		AvatarUrl: u.AvatarUrl,
	// 		CreatedAt: u.CreatedAt,
	// 		DeletedAt: u.DeletedAt,
	// 		UpdatedAt: u.UpdatedAt,
	// 	})
	// }
	return res, total, nil
}

func GetUser(ctx context.Context, userId string) (*base.User, error) {
	var user *base.User
	if err := DB.WithContext(ctx).Model(model.User{}).Where("user_id=?", userId).Find(&user); err != nil {
		logrus.Info(err)
		return user, errors.Wrapf(err.Error, "GetUser failed,err:%v", err)
	}
	return user, nil
}

func UserExist(ctx context.Context, uid string) ([]*base.User, error) {
	var user []*base.User
	if err := DB.WithContext(ctx).Model(&model.User{}).Where("user_id=?", uid).Find(&user); err != nil {
		return nil, errors.Wrapf(err.Error, "User not exist,err:%v", err)
	}
	return user, nil
}

func RemoveDuplicate(ctx context.Context, username string) (err error, flag bool) {
	var user base.User
	var count int64
	if err := DB.WithContext(ctx).Model(&user).Where("user_name=?", username).Count(&count).Error; err != nil {
		return errors.Wrap(err, "Query user failed!"), false
	}
	if count > 0 {
		return errors.Errorf("User is duplicate,please register again!"), false
	}
	return nil, true
}

func UploadAvatarUrl(ctx context.Context, uid, avatarUrl string) error {
	user, err := UserExist(ctx, uid)
	if err != nil || user == nil {
		return err
	}
	if err = DB.WithContext(ctx).Model(&base.User{}).Where("uid = ?").Update("avatar_url", avatarUrl).Error; err != nil {
		return err
	}
	return nil
}

func GetMaxId(ctx context.Context) (int64, error) {
	var maxId int64
	if err := DB.WithContext(ctx).Model(&base.User{}).Select("max(user_id)").Scan(&maxId).Error; err != nil {
		return 0, err
	}
	return maxId, nil
}
