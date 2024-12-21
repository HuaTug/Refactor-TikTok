package oss

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/minio/minio-go/v7"
)

var (
	minioClient *minio.Client
)

func UploadAvatar(data *[]byte, dataSize int64, uid string, tag string) (string, error) {
	// 在上传头像时需要满足 先删除旧的头像后 再上传新的头像
	deleteAvatar(uid)
	bucketName := "picture"
	var suffix string
	switch tag {
	case "image/jpeg", "image/jpg":
		suffix = ".jpg"
	case "image/png":
		suffix = ".png"
	default:
		return "", fmt.Errorf("unsupported image format: %s", tag)
	}

	objectName := "avatar/" + uid + suffix

	_, err := minioClient.PutObject(context.Background(), bucketName, objectName, bytes.NewReader(*data), dataSize, minio.PutObjectOptions{ContentType: tag})
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("http://%s/%s/%s", "localhost:9091/browser", bucketName, objectName), nil
}

func deleteAvatar(uid string) {
	bucketName := "picture"
	keys := []string{
		"avatar/" + uid + ".jpg",
		"avatar/" + uid + ".jpeg",
		"avatar/" + uid + ".png",
	}
	for _, key := range keys {
		err := minioClient.RemoveObject(context.Background(), bucketName, key, minio.RemoveObjectOptions{})
		if err != nil {
			log.Printf("Failed to delete %s: %v", key, err)
		}
	}
}

func UploadVideo(path, vid string) (string, error) {
	// 这是我上传视频文件的路径
	objectName := "video/" + vid + "/video.mp4"
	bucketName := "video"

	// 上传视频文件到 MinIO
	_, err := minioClient.FPutObject(context.Background(), bucketName, objectName, path, minio.PutObjectOptions{ContentType: "video/mp4"})
	if err != nil {
		hlog.Info(err)
		return "", err
	}

	// 返回视频的 URL
	return fmt.Sprintf("http://%s/%s/%s", "localhost:9091/browser", bucketName, objectName), nil
}

//ToDo: 上传视频封面时 也可以让用户自定义封面
func UploadVideoCover(path, vid string) (string, error) {
	objectName := "picture/" + vid + "/cover.jpg"
	bucketName := "picture"

	// 上传封面文件到 MinIO
	_, err := minioClient.FPutObject(context.Background(), bucketName, objectName, path, minio.PutObjectOptions{ContentType: "image/jpeg"})
	if err != nil {
		return "", err
	}

	// 返回封面的 URL
	return fmt.Sprintf("http://%s/%s/%s", "localhost:9091/browser", bucketName, objectName), nil
}

// 生成了下载的预签名地址，然后下载视频
func GeneratePreUrl(bucketName, objectName, vid string) (string, error) {
	presignedURL, err := minioClient.PresignedGetObject(context.Background(), bucketName, objectName, 30*time.Minute, nil)
	if err != nil {
		hlog.Info(err)
		return "", err
	}
	hlog.Info("Download URL:", presignedURL)
	
	// 发起 GET 请求下载视频
	resp, err := http.Get(presignedURL.String())
	if err != nil {
		hlog.Info(err)
	}
	defer resp.Body.Close()

	outputDir := "../../Download_video/"

	// 创建文件夹
	err = os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		hlog.Info("Error creating directory: ", err)
		return "", err
	}

	// 创建文件
	filePath := outputDir + "videos" + vid + ".mp4"

	out, err := os.Create(filePath)
	if err != nil {
		hlog.Info(err)
	}
	defer out.Close()

	// 将下载的内容写入文件
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		hlog.Info("Error saving video: ", err)
	}

	hlog.Info("Video Downloaded successfully!")

	return presignedURL.String(), nil
}
