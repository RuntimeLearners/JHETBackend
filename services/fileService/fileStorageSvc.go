package fileservice

import (
	configreader "JHETBackend/configs/configReader"
	"crypto/rand"
	"io"
	"log"
	"math/big"
	"os"
)

// #####PUBLIC#####

// 统一从io读文件存盘操作
func SaveUploadedFile(ior *io.Reader) (string, error) {
	dir := configreader.GetConfig().FileObject.Dir

	fileName := randStrGenerater(32)
	filePath := dir + "/" + fileName
	dst, err := os.Create(filePath)

	if err != nil {
		return "", err
	}
	defer dst.Close()
	_, err = io.Copy(dst, *ior)
	if err != nil {
		return "", err
	}
	log.Printf("[INFO][FileCtrl] New file uploaded, file: %v", dst.Name())
	return fileName, nil
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
