package oss

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/minio/minio-go/v7"
)

var (
	minioClient *minio.Client
	bucketName  = "picture"
)

func UploadAvatar(data *[]byte, dataSize int64, uid string, tag string) (string, error) {
	// 在上传头像时需要满足 先删除旧的头像后 再上传新的头像
	deleteAvatar(uid)

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

func UploadVideoCover(path, vid string) (string, error) {
	objectName := "video/" + vid + "/cover.jpg"
	bucketName := "picture"

	// 上传封面文件到 MinIO
	_, err := minioClient.FPutObject(context.Background(), bucketName, objectName, path, minio.PutObjectOptions{ContentType: "image/jpeg"})
	if err != nil {
		return "", err
	}

	// 返回封面的 URL
	return fmt.Sprintf("http://%s/%s/%s", "localhost:9091/browser", bucketName, objectName), nil
}
