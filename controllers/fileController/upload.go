package filectrl

import (
	"JHETBackend/common/exception"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

// FileController 可以把一组文件相关的 handler 挂在这里
type FileController struct {
	// 如果后面要加数据库、日志、配置，可以丢进来
	UploadDir string // 文件落盘目录
}

// NewFileController 构造函数，外部可以注入目录
func NewFileController(uploadDir string) *FileController {
	return &FileController{UploadDir: uploadDir}
}

// UploadFile 处理单文件上传  POST /upload
func (fc *FileController) UploadFile(c *gin.Context) {
	// 1. 取出文件（multipart）
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.Error(exception.ApiNoFormFile)
		return
	}
	defer file.Close()

	// 2. 简单校验
	const maxSize = 10 << 20 // 10 MB
	if header.Size > maxSize {
		c.JSON(http.StatusRequestEntityTooLarge, gin.H{"msg": "文件太大"})
		return
	}
	ext := filepath.Ext(header.Filename)
	allow := map[string]bool{".jpg": true, ".jpeg": true, ".png": true}
	if !allow[ext] {
		c.Error(exception.ApiFileNotSupported)
		return
	}

	//TODO: 下面这些AI代码简直是依托，还要改 MucheXD/09.21

	// 3. 生成唯一文件名
	newName := fmt.Sprintf("%d_%s", time.Now().Unix(), ext)
	dstPath := filepath.Join(fc.UploadDir, newName)

	// 4. 落盘
	if err := c.SaveUploadedFile(header, dstPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "保存失败"})
		return
	}

	// 5. 返回访问地址（这里简单拼一个相对路径，生产环境最好走 CDN/对象存储）
	url := "/static/" + newName
	c.JSON(http.StatusOK, gin.H{"url": url})
}
