package utils

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"math/big"
	"net/smtp"
	"regexp"
	"strings"
)

const (
	// 定义字符集
	chars  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length = 6
)

// generateRandomString 生成指定长度的乱序字符串
func generateRandomString() (string, error) {
	var sb strings.Builder
	sb.Grow(length)

	// 使用 crypto/rand 生成随机字符
	for i := 0; i < length; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return "", err
		}
		sb.WriteByte(chars[index.Int64()])
	}

	return sb.String(), nil
}

func IsValidEmail(email string) bool {
	// 定义正则表达式以匹配特定的电子邮件格式
	const emailRegex = `^[0-9]{10}@stumail\.nwu\.edu\.cn$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func SendEmail(to string) (string, error) {
	smtpHost := "smtp.163.com"
	smtpPort := "587"
	smtpUser := "xu_zh0105@163.com"
	smtpPassword := "LY3gCRfaVs7yHY3x"
	// if !IsValidEmail(to) {
	// 	return "", fmt.Errorf("invalid email address,please check your email address")
	// }
	// 构建地址
	addr := smtpHost + ":" + smtpPort

	// 创建 TLS 配置
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpHost,
	}

	// 连接到 SMTP 服务器
	conn, err := tls.Dial("tcp", addr, tlsconfig)
	if err != nil {
		return "", fmt.Errorf("failed to connect to server: %w", err)
	}
	c, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		return "", fmt.Errorf("failed to create SMTP client: %w", err)
	}

	// 进行身份验证
	if err = c.Auth(smtp.PlainAuth("", smtpUser, smtpPassword, smtpHost)); err != nil {
		return "", fmt.Errorf("authentication failed: %w", err)
	}

	// 创建邮件内容
	verificationCode, err := generateRandomString()
	if err != nil {
		return "", fmt.Errorf("failed to generate verification code: %w", err)
	}
	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + "欢迎西北大学的用户登录!" + "\r\n" +
		"\r\n" +
		verificationCode)

	// 发送邮件
	if err = c.Mail(smtpUser); err != nil {
		return "", fmt.Errorf("failed to set sender: %w", err)
	}
	if err = c.Rcpt(to); err != nil {
		return "", fmt.Errorf("failed to set recipient: %w", err)
	}

	w, err := c.Data()
	if err != nil {
		return "", fmt.Errorf("failed to write message: %w", err)
	}
	_, err = w.Write(msg)
	if err != nil {
		return "", fmt.Errorf("failed to get data writer: %w", err)
	}
	err = w.Close()
	if err != nil {
		return "", fmt.Errorf("failed to close data writer: %w", err)
	}
	// 关闭连接
	c.Quit()
	return verificationCode, nil
}
