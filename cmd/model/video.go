package model

type Video struct {
	VideoId        int64
	UserId         int64
	VideoUrl       string
	CoverUrl       string
	Title          string
	Description    string
	VisitCount     int64
	LikeCount      int64
	CommentCount   int64
	CreatedAt      string
	UpdatedAt      string
	DeletedAt      string
	Open           int64
	AuditStatus    int64
	ShareCount     string
	LabelNames     string
	Duration       string
	FavoritesCount string
	HistoryCount   string
}

// 收藏夹
type Favorite struct {
	FavoriteId  int64
	UserId      int64
	Name        string
	Description string
	CoverUrl    string
	CreatedAt   string
	DeletedAt   string
}

// 收藏夹中的视频
type FavoritesVideos struct {
	FavoriteVideoId int64
	FavoriteId      int64
	VideoId         int64
	UserId          int64
}

type VideoShare struct {
	VideoShareId int64
	UserId       int64
	VideoId      int64
	ToUserId     int64
	CreatedAt    string
	DeletedAt    string
}

type UserVideoWatchHistory struct {
	UserVideoWatchHistoryId int64
	UserId                  int64
	VideoId                 int64
	WatchTime               string
	DeletedAt               string
}
