package filecontroller

import (
	"JHETBackend/common/exception"
	accountControllers "JHETBackend/controllers/accountControllers"
	fileservice "JHETBackend/services/fileService"
	"JHETBackend/services/userService"
	"encoding/hex"
	"io"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

// TODO: 文件上传的逻辑待整理 现有代码有混乱和重复的部分

// 上传用户文件
// 在使用这个控制器前，请先检查用户权限，以免接口被滥用
func UploadAttachment(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.Error(exception.ApiNoFormFile)
		return
	}
	if fileHeader.Size > int64(10485760) { // 对附件文件限制 10MB
		c.Error(exception.ApiFileTooLarge)
	}
	fileHandler, err := getFileHandler(fileHeader)
	if err != nil {
		c.Error(err) // 由于 getFileHandler 也使用统一错误，因此可以直接返回
		return
	}

	// 取出前端传入的SHA256
	cSha256Hex := c.PostForm("SHA256")
	cSha256Data, err := hex.DecodeString(cSha256Hex)
	if (len(cSha256Hex) != 64) || (err != nil) {
		cSha256Data = nil
	}

	uuid, err := fileservice.SaveUploadedFile(&fileHandler, cSha256Data)

	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, gin.H{
		"file_uuid": uuid,
	})
}

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
	accountID, err := accountControllers.GetAccountIDFromContext(c)
	if err != nil {
		c.Error(err)
		return
	}
	userService.UpdateAvatar(accountID, fileHandler)
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
