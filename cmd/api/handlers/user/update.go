package handlers

import (
	"context"
	"io"

	"HuaTug.com/cmd/api/rpc"
	"HuaTug.com/kitex_gen/users"
	jwt "HuaTug.com/pkg"
	"HuaTug.com/pkg/errno"
	"HuaTug.com/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"
)

func UpdateUser(ctx context.Context, c *app.RequestContext) {

	var v interface{}
	var err error
	if v, err = jwt.ConvertJWTPayloadToString(ctx, c); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	}
	userId := utils.Transfer(v)
	var update UpdateParam
	if err := c.Bind(&update); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	password, _ := utils.Crypt(update.PassWord)

	// 获取上传的文件
	uploadData, err := c.FormFile("file")
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	//	打开文件流
	file, err := uploadData.Open()
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	defer file.Close()
	// 处理文件内容，将文件内容读取到内存中
	fileContent, err := io.ReadAll(file)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	resp, err := rpc.UpdateUser(ctx, &users.UpdateUserRequest{
		UserId:   userId,
		UserName: update.UserName,
		Password: password,
		Data:     fileContent,
		Filesize: uploadData.Size,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	
	SendResponse(c, errno.Success, resp)
}
