package filecontroller

import (
	"JHETBackend/common/exception"
	configreader "JHETBackend/configs/configReader"
	"JHETBackend/services/userService"
	"io"
	"mime/multipart"
	"sync"

	//"crypto/md5" hash算法库

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

func UploadAvatar(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.Error(exception.ApiNoFormFile)
		return
	}
	if fileHeader.Size > int64(102400) { // 对头像文件限制 100kb
		c.Error(exception.ApiFileTooLarge)
	}
	fileHandler, err := getFileHandler(fileHeader)
	if err != nil {
		c.Error(err) // 由于 getFileHandler 也使用统一错误，因此可以直接返回
		return
	}
	userService.UploadAvatar(fileHandler)
}

func getFileHandler(fileHeader *multipart.FileHeader) (io.Reader, error) {
	initOnce.Do(initFileController)
	// 打开文件
	fileHandler, err := fileHeader.Open()
	if err != nil {
		return nil, exception.ApiFileCannotOpen
	}

	defer fileHandler.Close() // 返回时关闭文件

	return fileHandler, nil
}
