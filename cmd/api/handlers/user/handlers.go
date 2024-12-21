// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

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

// 在Go语言中，结构体标签的格式应该是 key:"value"，并且多个键值对之间应该用空格分隔。
type UserParam struct {
	UserName string `form:"user_name" json:"username"`
	PassWord string `form:"password"  json:"password"`
	Email    string `form:"email"`
	Sex      int64  `form:"sex"`
}

type LoginParam struct {
	UserName string `form:"user_name" json:"username"`
	PassWord string `form:"password"  json:"password"`
	Email    string `form:"email"`
}

type QueryParam struct {
	PageNum  int64  `form:"page_num"`
	PageSize int64  `form:"page_size"`
	Keyword  string `form:"keyword"`
}

type UpdateParam struct {
	UserName string `form:"user_name" `
	PassWord string `form:"password"`
}

type SendCodeParam struct {
	Email string `form:"email"`
}

type VerifyCodeParam struct {
	Email string `form:"email"`
	Code  string `form:"code"`
}
