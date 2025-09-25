package fileservice

import (
	"crypto/rand"
	"io"
)


func SaveUploadedFile(ior *io.Reader) (dst string, err error) {
	if err := ior.Read(make([]byte, 1)); err != nil && err != io.EOF {
}

// #####PRIVATE#####

// 生成随机字符串 用于临时文件名
func randStrGenerater(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}