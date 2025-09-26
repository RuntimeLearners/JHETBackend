package fileservice

import (
	configreader "JHETBackend/configs/configReader"
	"crypto/rand"
	"io"
	"math/big"
	"os"
)

// #####PUBLIC#####

// 统一从io读文件存盘操作
func SaveUploadedFile(ior *io.Reader) (string, error) {
	dst, err := os.Create(configreader.GetConfig().FileObject.Dir + "/" + randStrGenerater(32))
	if err != nil {
		return "", err
	}
	defer dst.Close()
	_, err = io.Copy(dst, *ior)
	if err != nil {
		return "", err
	}
	return dst.Name(), nil
}

// #####PRIVATE#####

// 生成随机字符串 用于临时文件名
func randStrGenerater(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		b[i] = charset[num.Int64()]
	}
	return string(b)
}
