package handlers

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"HuaTug.com/cmd/api/rpc"
	"HuaTug.com/cmd/video/infras/redis"

	"HuaTug.com/kitex_gen/videos"
	jwt "HuaTug.com/pkg"
	"HuaTug.com/pkg/errno"
	"HuaTug.com/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func VideoPublishUploading(ctx context.Context, c *app.RequestContext) {
	var err error
	var v interface{}
	var UserId int64
	var VideoPublish VideoPublishUploadingParam
	if err = c.BindAndValidate(&VideoPublish); err != nil {
		hlog.Info(err)
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	if v, err = jwt.ConvertJWTPayloadToString(ctx, c); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	} else {
		UserId = utils.Transfer(v)
	}
	Data, err := c.FormFile("data")
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	fileSize := Data.Size
	res, err := redis.GetChunkInfo(fmt.Sprint(UserId), VideoPublish.Uuid)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	value, _ := strconv.ParseInt(res[0], 10, 64)
	chunksize := int(fileSize) / int(value)
	offset := (int(VideoPublish.ChunkNumber) - 1) * chunksize
	file, err := Data.Open()
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	defer file.Close()

	_, err = file.Seek(int64(offset), io.SeekStart)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	}
	buffer := make([]byte, chunksize)
	bytesRead, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	chunkFileName := fmt.Sprintf("%s_part_%d", VideoPublish.FileName, VideoPublish.ChunkNumber)

	md5hash := utils.GetBytesMD5(buffer[:bytesRead])
	resp, err := rpc.VideoPublishUploading(ctx, &videos.VideoPublishUploadingRequest{
		UserId:      UserId,
		Data:        buffer[:bytesRead],
		Filename:    chunkFileName,
		IsM3u8:      VideoPublish.Is_M3U8,
		Md5:         md5hash,
		Uuid:        VideoPublish.Uuid,
		ChunkNumber: VideoPublish.ChunkNumber,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, resp)
}
