package utils

import (
	"strings"

	"HuaTug.com/config"
)

func GetMysqlDsn() string {
	//生成数据库的dsn
	dsn := strings.Join([]string{config.ConfigInfo.Mysql.Username, ":",
		config.ConfigInfo.Mysql.Password, "@tcp(", config.ConfigInfo.Mysql.Addr, ")/",
		config.ConfigInfo.Mysql.Database, "?charset=" + config.ConfigInfo.Mysql.Charset + "&parseTime=true"}, "") //nolint:lll

	return dsn
}
