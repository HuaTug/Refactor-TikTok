// Code generated by hertz generator. DO NOT EDIT.

package interactions

import (
	interactions "HuaTug.com/cmd/api/handlers/interaction"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_v1 := root.Group("/v1", _v1Mw()...)
		{
			_action := _v1.Group("/action", _actionMw()...)
			_action.POST("/like", append(_likeactionMw(), interactions.LikeAction)...)
			_action.GET("/list", append(_likelistMw(), interactions.LikeList)...)
		}
		{
			_comment := _v1.Group("/comment", _commentMw()...)
			_comment.DELETE("/delete", append(_deletecommentMw(), interactions.DeleteComment)...)
			_comment.GET("/list", append(_listcommentMw(), interactions.ListComment)...)
			_comment.POST("/publish", append(_createcommentMw(), interactions.CreateComment)...)
		}
	}
}