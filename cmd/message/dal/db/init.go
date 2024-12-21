package db

import (
	"HuaTug.com/config"
	"HuaTug.com/pkg/sharding"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

var (
	DB *gorm.DB
)
func Load() {
	var err error
	//dsn := utils.GetMysqlDsn()
	
	hlog.Info(config.ConfigInfo.Mysql.Addr)
	dsn := config.ConfigInfo.Mysql.Username + ":" + config.ConfigInfo.Mysql.Password + "@tcp(" + config.ConfigInfo.Mysql.Addr + ")/" + config.ConfigInfo.Mysql.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	if err = DB.Use(gormopentracing.New()); err != nil {
		panic(err)
	}
	if err = DB.Use(sharding.NewSharding("to_user_id", 4, "messages")); err != nil {
		panic(err)
	}

	hlog.Infof("Mysql connected successfully.")
}
