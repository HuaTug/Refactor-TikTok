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

type CreateCommentParam struct {
	VideoId   int64  `form:"video_id"`
	CommentId int64  `form:"comment_id"`
	Mode      int64  `form:"mode"`
	Content   string `form:"content"`
}

type ListCommentParam struct {
	VideoId   int64 `form:"video_id"`
	CommentId int64 `form:"comment_id"`
	PageNum   int64 `form:"page_num"`
	PageSize  int64 `form:"page_size"`
}

type DeleteCommentParam struct {
	VideoId    int64 `form:"video_id"`
	CommentId  int64 `form:"comment_id"`
	FromUserId int64 `form:"from_user_id"`
}

type LikeParam struct {
	VideoId    int64  `form:"video_id"`
	CommentId  int64  `form:"comment_id"`
	ActionType string `form:"action_type"`
}

type LikeListParam struct {
	PageNum  int64 `form:"page_num"`
	PageSize int64 `form:"page_size"`
}
