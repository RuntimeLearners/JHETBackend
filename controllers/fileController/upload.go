package filecontroller

import (
	"JHETBackend/common/exception"
	configreader "JHETBackend/configs/configReader"
	"math/rand"
	"net/http"
	"path/filepath"
	"sync"

	"github.com/gin-gonic/gin"
)

var fileSaveDir string
var largeFileSize int

// 初始化模块 载入保存目录等
func initFileController() {
	fileSaveDir = configreader.GetConfig().FileObject.Dir
	largeFileSize = configreader.GetConfig().FileObject.LargeFileSize
}

// UploadFile 处理单文件上传  POST /upload

var initOnce sync.Once

func UploadFile(c *gin.Context) {
	initOnce.Do(initFileController)

	// 取出文件（multipart）
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.Error(exception.ApiNoFormFile)
		return
	}
	// 函数结束时关闭文件
	defer func() {
		file.Close()
		// TODO：将文件名改为哈希，方便后期校验
	}()

	// 校验文件大小 之后可以在权限列表中增加大文件上传权限(ENHANCEMENT)
	if header.Size > int64(largeFileSize) {
		c.Error(exception.ApiFileTooLarge)
		return
	}

	// 生成临时文件名 32长度随机字符确保随机性
	tmpName := "tmp_" + randStrGenerater(32)
	// 拼接路径
	dstPath := filepath.Join(fileSaveDir, tmpName)

	// 文件落盘
	if err := c.SaveUploadedFile(header, dstPath); err != nil {
		c.JSON(http.StatusInternalServerError, exception.FileCannotSaveUploaded)
		return
	}

	// 返回文件名交还前端
	// TODO 确定API规则后填写此处

}

func randStrGenerater(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
