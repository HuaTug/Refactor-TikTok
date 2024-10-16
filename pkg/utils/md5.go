package utils

import (
	"crypto/md5" //nolint:gosec
	"encoding/hex"
	"io"
	"os"
)

func MD5(str string) string {
	h := md5.New() //nolint:gosec
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func GetFileMD5(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return ``, err
	}
	defer file.Close()

	md5Hash := md5.New()
	if _, err := io.Copy(md5Hash, file); err != nil {
		return ``, err
	}

	return hex.EncodeToString(md5Hash.Sum(nil)), nil
}

func GetBytesMD5(data []byte) string {
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}
