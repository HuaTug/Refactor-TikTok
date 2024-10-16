package model

type Follow struct {
	FollowId    int64
	FollowingId int64
	FollowersId int64
	CreatedAt   string
	DeletedAt   string
}
