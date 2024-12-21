package utils

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func M3u8ToMp4(input, output string) error {
	err := ffmpeg.Input(input).
		Output(output, ffmpeg.KwArgs{
			"c:v":     "copy",
			"absf":    "aac_adtstoasc",
			"b:v":     "4000k",
			"bufsize": "4000k",
		}).
		OverWriteOutput().
		Run()
	if err != nil {
		return err
	}
	return nil
}

func GetVideoThumnail(videoPath, outputDir string) (string, error) {
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return "", errors.WithMessage(err, "Failed to create folders")
	}
	outputPath := filepath.Join(outputDir, "thumnail.jpg")
	err := ffmpeg.Input(videoPath).
		Output(outputPath, ffmpeg.KwArgs{
			"ss":      "00:00:00",
			"vframes": "1",
		}).
		OverWriteOutput().
		Run()
	if err != nil {
		return "", errors.WithMessage(err, "Failed to generate the thumnail")
	}
	return outputPath, nil
}
