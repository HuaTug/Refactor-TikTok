package handlers

import (
	"HuaTug.com/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SendResponse pack response
func SendResponse(c *app.RequestContext, err error, data interface{}) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, Response{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		Data:    data,
	})
}

type UpLoadvideoParam struct {
	ContentType string `json:"content_type" form:"content_type"`
	ObjectName  string `json:"object_name" form:"object_name"`
	BucketName  string `json:"bucket_name" form:"bucket_name"`
	Title       string `json:"title" form:"title"`
	CoverUrl    string `json:"cover_url" form:"cover_url"`
}

type FeedListParam struct {
	LastTime string `json:"last_time" form:"last_time"`
}

type VideoFeedListParam struct {
	AuthorId int64 `form:"author_id" `
	PageNum  int64 `form:"page_num"`
	PageSize int64 `form:"page_size"`
}

type VideoSearchParam struct {
	Keyword  string `form:"keyword"`
	PageNum  int64  `form:"page_num"`
	PageSize int64  `form:"page_size"`
	FromDate string `form:"from_date"`
	ToDate   string `form:"to_date"`
}

type VideoPublishStartParam struct {
	Title            string `form:"title"`
	Description      string `form:"description"`
	LabName          string `form:"lab_name"`
	Category         string `form:"category"`
	Open             int64  `form:"open"`
	ChunkTotalNumber int64  `form:"chunk_total_number"`
}

type VideoPublishUploadingParam struct {
	Uuid        string `form:"uuid"`
	Data        byte   `form:"data"`
	Is_M3U8     bool   `form:"is_m3u8"`
	FileName    string `form:"filename"`
	ChunkNumber int64  `form:"chunk_number"`
}

type VideoPublishCompleteParam struct {
	Uuid string `form:"uuid"`
}

type VideoPublishCancleParam struct {
	Uuid string `form:"uuid"`
}

type VideoDeleteParam struct {
	VideoId int64 `form:"video_id"`
}

type VideoStreamParam struct {
	Index string `form:"index"`
}

type CreateFavoriteParam struct {
	UserId      int64  `form:"user_id"`
	Name        string `form:"name"`
	Description string `form:"description"`
	CoverUrl    string `form:"cover_url"`
}

type GetFavoriteListParam struct {
	PageNum  int64 `form:"page_num"`
	PageSize int64 `form:"page_size"`
}

type AddFavoriteVideoParam struct {
	FavoriteId int64 `form:"favorite_id"`
	VideoId    int64 `form:"video_id"`
}

type GetFavoriteVideoListParam struct {
	FavoriteId int64 `form:"favorite_id"`
	PageNum    int64 `form:"page_num"`
	PageSize   int64 `form:"page_size"`
}

type GetFavoriteVideoByIdParam struct {
	FavoriteId int64 `form:"favorite_id"`
	VideoId    int64 `form:"video_id"`
}

type DeleteFavoriteParam struct {
	FavoriteId int64 `form:"favorite_id"`
}

type DeleteVideoFromFavoriteParam struct {
	FavoriteId int64 `form:"favorite_id"`
	VideoId    int64 `form:"video_id"`
}
type GetVideoFromFavoriteParam struct {
	FavoriteId int64 `form:"favorite_id"`
	VideoId    int64 `form:"video_id"`
	PageNum    int64 `form:"page_num"`
	PageSize   int64 `form:"page_size"`
}

type SharedVideoParam struct {
	VideoId  int64 `form:"video_id"`
	ToUserId int64 `form:"to_user_id"`
}
