`项目结构`

```
├── biz
│   ├── config
│   │   └── mysql
│   └── dal
│       └── cache
│           └── redis
│               └── redis.conf
├── cache
│   ├── comment.go
│   ├── init.go
│   ├── redis
│   │   └── redis.conf
│   ├── relation.go
│   ├── user.go
│   └── video.go
├── cmd
│   ├── api
│   │   ├── api
│   │   ├── handlers
│   │   │   ├── comment
│   │   │   │   ├── createcomment.go
│   │   │   │   ├── deletecomment.go
│   │   │   │   ├── handlers.go
│   │   │   │   └── listcomment.go
│   │   │   ├── favorite
│   │   │   │   ├── favorite.go
│   │   │   │   ├── handlers.go
│   │   │   │   └── listfavorite.go
│   │   │   ├── relation
│   │   │   │   ├── handlers.go
│   │   │   │   ├── relation.go
│   │   │   │   └── relationpage.go
│   │   │   ├── user
│   │   │   │   ├── delete.go
│   │   │   │   ├── getuserinfo.go
│   │   │   │   ├── handlers.go
│   │   │   │   ├── login.go
│   │   │   │   ├── query.go
│   │   │   │   ├── registers.go
│   │   │   │   └── update.go
│   │   │   └── video
│   │   │       ├── feedlist.go
│   │   │       ├── handlers.go
│   │   │       ├── uploadvideo.go
│   │   │       ├── videolist.go
│   │   │       ├── videopopular.go
│   │   │       └── videosearch.go
│   │   ├── main.go
│   │   ├── router
│   │   │   ├── authfunc
│   │   │   │   └── auth_logic.go
│   │   │   ├── comment
│   │   │   │   ├── comment.go
│   │   │   │   └── middleware.go
│   │   │   ├── favorite
│   │   │   │   ├── favorite.go
│   │   │   │   └── middleware.go
│   │   │   ├── publish
│   │   │   │   ├── middleware.go
│   │   │   │   └── publish.go
│   │   │   ├── register.go
│   │   │   ├── relation
│   │   │   │   ├── middleware.go
│   │   │   │   └── relation.go
│   │   │   ├── user
│   │   │   │   ├── middleware.go
│   │   │   │   └── user.go
│   │   │   └── video
│   │   │       ├── middleware.go
│   │   │       └── video.go
│   │   ├── router_gen.go
│   │   ├── router.go
│   │   └── rpc
│   │       ├── comment.go
│   │       ├── favorite.go
│   │       ├── init.go
│   │       ├── publish.go
│   │       ├── relation.go
│   │       ├── user.go
│   │       └── video.go
│   ├── comment
│   │   ├── dal
│   │   │   ├── db
│   │   │   │   ├── comment.go
│   │   │   │   └── init.go
│   │   │   └── init.go
│   │   ├── handler.go
│   │   ├── main.go
│   │   └── service
│   │       ├── createcommentservice.go
│   │       ├── deletecommentservice.go
│   │       └── listcommentservice.go
│   ├── favorite
│   │   ├── dal
│   │   │   ├── db
│   │   │   │   ├── favorite.go
│   │   │   │   └── init.go
│   │   │   └── init.go
│   │   ├── favorite
│   │   ├── handler.go
│   │   ├── main.go
│   │   └── service
│   │       ├── favorite.go
│   │       ├── listfavorite.go
│   │       └── unfavorite.go
│   ├── mw
│   │   ├── Es
│   │   │   ├── init.go
│   │   │   └── video_index_init.go
│   │   ├── jwt.go
│   │   └── sentinels
│   │       └── Flow.go
│   ├── publish
│   │   ├── dal
│   │   │   ├── db
│   │   │   │   ├── init.go
│   │   │   │   └── publish.go
│   │   │   └── init.go
│   │   ├── handler.go
│   │   ├── main.go
│   │   ├── publish
│   │   └── service
│   │       ├── uploadvideo.go
│   │       └── videocreate.go
│   ├── relation
│   │   ├── dal
│   │   │   ├── db
│   │   │   │   ├── init.go
│   │   │   │   └── relaiton.go
│   │   │   └── init.go
│   │   ├── handler.go
│   │   ├── main.go
│   │   ├── relation
│   │   └── service
│   │       ├── following.go
│   │       └── followlist.go
│   ├── user
│   │   ├── dal
│   │   │   ├── db
│   │   │   │   ├── init.go
│   │   │   │   └── user.go
│   │   │   └── init.go
│   │   ├── handler.go
│   │   ├── main.go
│   │   ├── service
│   │   │   ├── createuser.go
│   │   │   ├── deleteuser.go
│   │   │   ├── getuserinfo.go
│   │   │   ├── loginuser.go
│   │   │   ├── query.go
│   │   │   └── updateuser.go
│   │   └── user
│   └── video
│       ├── dal
│       │   ├── db
│       │   │   ├── init.go
│       │   │   └── video.go
│       │   └── init.go
│       ├── handler.go
│       ├── main.go
│       ├── service
│       │   ├── feedlist.go
│       │   ├── videolist.go
│       │   ├── videopopular.go
│       │   └── videosearch.go
│       └── video
├── config
│   ├── config.go
│   ├── config.yml
│   ├── mysql
│   │   └── mysql.sql
│   └── types.go
├── docker-compose.yml
├── golangci-lint
│   └── 2024-5-8.txt
├── go.mod
├── go.sum
├── idl
│   ├── chat.thrift
│   ├── comment.thrift
│   ├── favorite.thrift
│   ├── publish.thrift
│   ├── relation.thrift
│   ├── user.thrift
│   └── video.thrift
├── kitex_gen
│   ├── comments
│   │   ├── comment.go
│   │   ├── commentservice
│   │   │   ├── client.go
│   │   │   ├── commentservice.go
│   │   │   ├── invoker.go
│   │   │   └── server.go
│   │   ├── k-comment.go
│   │   └── k-consts.go
│   ├── favorites
│   │   ├── favorite.go
│   │   ├── favoriteservice
│   │   │   ├── client.go
│   │   │   ├── favoriteservice.go
│   │   │   ├── invoker.go
│   │   │   └── server.go
│   │   ├── k-consts.go
│   │   └── k-favorite.go
│   ├── publishs
│   │   ├── k-consts.go
│   │   ├── k-publish.go
│   │   ├── publish.go
│   │   └── uploadvideoservice
│   │       ├── client.go
│   │       ├── invoker.go
│   │       ├── server.go
│   │       └── uploadvideoservice.go
│   ├── relations
│   │   ├── followservice
│   │   │   ├── client.go
│   │   │   ├── followservice.go
│   │   │   ├── invoker.go
│   │   │   └── server.go
│   │   ├── k-consts.go
│   │   ├── k-relation.go
│   │   └── relation.go
│   ├── users
│   │   ├── k-consts.go
│   │   ├── k-user.go
│   │   ├── user.go
│   │   └── userservice
│   │       ├── client.go
│   │       ├── invoker.go
│   │       ├── server.go
│   │       └── userservice.go
│   └── videos
│       ├── k-consts.go
│       ├── k-video.go
│       ├── video.go
│       └── videoservice
│           ├── client.go
│           ├── invoker.go
│           ├── server.go
│           └── videoservice.go
├── kitex_gen.sh
├── Makefile
├── pkg
│   ├── bound
│   │   └── cpu.go
│   ├── constants
│   │   ├── constant.go
│   │   └── util.go
│   ├── errno
│   │   └── errno.go
│   ├── middleware
│   │   ├── client.go
│   │   ├── common.go
│   │   └── server.go
│   ├── tracer
│   │   └── tracer.go
│   └── utils
│       ├── crypto.go
│       ├── dsn.go
│       ├── md5.go
│       └── transfer.go
└── README.md
```