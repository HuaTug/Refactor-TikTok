package model

type User struct {
	UserId    int64  `json:"user_id"`
	UserName  string `json:"user_name"`
	AvatarUrl string `json:"avatar_url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}
