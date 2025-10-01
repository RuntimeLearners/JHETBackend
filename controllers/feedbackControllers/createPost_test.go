package feedbackcontrollers_test

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	feedbackcontrollers "JHETBackend/controllers/feedbackControllers"
	filecontroller "JHETBackend/controllers/fileController"

	"github.com/gin-gonic/gin"
)

func Test_CreateNewPost(t *testing.T) {

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// 注册上传和发帖路由
	router.POST("/upload", filecontroller.UploadAttachment)
	router.POST("/feedback",
		func(c *gin.Context) {
			// 模拟中间件设置AccountID
			c.Set("AccountID", uint64(14))
			filecontroller.UpdateAvatar(c)
		}, feedbackcontrollers.CreateFeedbackPost)

	// 1. 上传文件
	filePath := "../../testfile/test_avatar.jpg"
	file, _ := os.Open(filePath)
	defer file.Close()

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	part, _ := writer.CreateFormFile("file", filepath.Base(filePath))
	_, _ = io.Copy(part, file)
	writer.Close()

	req := httptest.NewRequest("POST", "/upload", &buf)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var uploadResp struct {
		FileUUID string `json:"file_uuid"`
	}
	_ = json.Unmarshal(w.Body.Bytes(), &uploadResp)

	// 2. 创建帖子
	postBody := map[string]interface{}{
		"title":            "测试标题",
		"urgency":          "medium",
		"content":          "测试内容",
		"isAnonymous":      false,
		"isPrivate":        false,
		"attachment_uuids": []string{uploadResp.FileUUID},
	}
	postBodyBytes, _ := json.Marshal(postBody)
	req2 := httptest.NewRequest("POST", "/feedback", bytes.NewReader(postBodyBytes))
	req2.Header.Set("Content-Type", "application/json")
	// 可根据需要设置认证信息，如cookie/header等

	w2 := httptest.NewRecorder()
	router.ServeHTTP(w2, req2)
	// 可根据业务进一步断言返回内容
}
