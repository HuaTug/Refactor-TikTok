package utils

import (
	"strings"

	"HuaTug.com/config"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func GetMysqlDsn() string {
	//生成数据库的dsn
	hlog.Info(config.ConfigInfo.Mysql.Username,",",config.ConfigInfo.Mysql.Password)
	dsn := strings.Join([]string{config.ConfigInfo.Mysql.Username, ":",
		config.ConfigInfo.Mysql.Password, "@tcp(", config.ConfigInfo.Mysql.Addr, ")/",
		config.ConfigInfo.Mysql.Database, "?charset=" + config.ConfigInfo.Mysql.Charset + "&parseTime=true"}, "") //nolint:lll

	return dsn
}
