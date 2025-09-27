package filecontroller_test

import (
	filecontroller "JHETBackend/controllers/fileController"
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_avatarUpload(t *testing.T) {

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/upload/avatar", func(c *gin.Context) {
		// 模拟中间件设置AccountID
		c.Set("AccountID", uint64(14))
		filecontroller.UpdateAvatar(c)
	})

	// 打开本地图片文件
	filePath := "../../testfile/test_avatar.jpg"
	file, err := os.Open(filePath)
	if err != nil {
		//cwd, _ := os.Getwd()
		//absPath := filepath.Join(cwd, filePath)
		absPath, _ := filepath.Abs(filePath)
		log.Println("完整路径:", absPath)
		t.Fatalf("无法打开测试图片文件: %v", err)
	}
	defer file.Close()

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	part, err := writer.CreateFormFile("file", "imagefile.png")
	if err != nil {
		t.Fatalf("创建表单文件失败: %v", err)
	}
	if _, err := io.Copy(part, file); err != nil {
		t.Fatalf("写入文件内容失败: %v", err)
	}
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/upload/avatar", &buf)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK && w.Code != http.StatusNoContent {
		t.Errorf("上传头像失败，HTTP状态码: %d, 响应: %s", w.Code, w.Body.String())
	}
}
