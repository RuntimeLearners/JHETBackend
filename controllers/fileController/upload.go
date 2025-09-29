package filecontroller

import (
	"JHETBackend/common/exception"
	accountcontroller "JHETBackend/controllers/accountControllers"
	"JHETBackend/services/userService"
	"io"
	"mime/multipart"

	//"crypto/md5" hash算法库 <<< 请使用sha256!(MucheXD)

	"github.com/gin-gonic/gin"
)

// UploadFile 处理单文件上传  POST /upload

// var initOnce sync.Once

func UpdateAvatar(c *gin.Context) {
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
	accountID, err := accountcontroller.GetAccountIDFromContext(c)
	if err != nil {
		c.Error(err)
		return
	}
	userService.UploadAvatar(accountID, fileHandler)
}

func getFileHandler(fileHeader *multipart.FileHeader) (io.Reader, error) {
	// initOnce.Do(initFileController)
	// 打开文件
	fileHandler, err := fileHeader.Open()
	if err != nil {
		return nil, exception.ApiFileCannotOpen
	}

	defer fileHandler.Close() // 返回时关闭文件

	return fileHandler, nil
}
