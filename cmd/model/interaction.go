package model

type Comment struct {
	CommentId int64
	UserId    int64
	VideoId   int64
	ParentId  int64
	Content   string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}

type CommentLike struct {
	CommentLikesId int64
	UserId         int64
	CommentId      int64
	CreatedAt      string
	DeletedAt      string
}

type VideoLike struct {
	VideoLikesId int64 
	UserId       int64
	VideoId      int64
	CreatedAt    string
	DeletedAt    string
}
