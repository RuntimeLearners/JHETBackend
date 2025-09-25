package fileservice

import (
	"crypto/rand"
	"io"
)

func SaveUploadedFile(ior *io.Reader) (dst string, err error) {
	// // 函数结束时关闭文件
	// defer func() {
	// 	fileContent.Close()
	// 	// TODO：将文件名改为哈希，方便后期校验
	// }()
	// // TODO: move to service

	// // 生成临时文件名 32长度随机字符确保随机性
	// tmpName := "tmp_" + randStrGenerater(32)
	// // 拼接路径
	// dstPath := filepath.Join(fileSaveDir, tmpName)

	// // 文件落盘
	// if err := c.SaveUploadedFile(fileHeader, dstPath); err != nil {
	// 	c.JSON(http.StatusInternalServerError, exception.FileCannotSaveUploaded)
	// 	return
	// }
}

// #####PRIVATE#####

// 生成随机字符串 用于临时文件名
func randStrGenerater(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		num, _ := rand.Int(rand.Reader, len(charset))
		b[i] = charset[num.Int64()]
	}
	return string(b)
}
