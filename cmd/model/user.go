package model

type User struct {
	UserId    int64  `json:"user_id"`
	UserName  string `json:"user_name"`
	AvatarUrl string `json:"avatar_url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type UserBehavior struct {
	UserBehaviorId int64  `json:"user_behavior_id"`
	UserId         int64  `json:"user_id"`
	VideoId        int64  `json:"video_id"`
	BehaviorType   string `json:"behavior_type"`
	BehaviorTime   string `json:"behavior_time"`
}
