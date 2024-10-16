package config

type config struct {
	Mysql    mysql
	Redis    redis
	Etcd     etcd
	RabbitMq rabbitmq
}

type mysql struct {
	Addr     string
	Database string
	Username string
	Password string
	Charset  string
}

type redis struct {
	Addr     string
	Password string
}
type etcd struct {
	Addr string
}
type rabbitmq struct {
	Addr     string
	Username string
	Password string
}
