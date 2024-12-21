package redis

type _Redis struct {
	Addr string
	DB   int
}

var (
	VideoInfo   = _Redis{Addr: "localhost:6379", DB: 1}
	VideoUpload = _Redis{Addr: "localhost:6379", DB: 2}
)
