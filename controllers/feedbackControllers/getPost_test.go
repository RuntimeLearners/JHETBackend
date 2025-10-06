package feedbackcontrollers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_GetFeedbackPosts(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/api/feedback", func(c *gin.Context) {
		getFeedbackPosts(c, false)
	})

	// 查询参数模板
	query := "?page=1&size=10"

	req, _ := http.NewRequest("GET", "/api/feedback"+query, nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	log.Printf("%v", w.Body.String())

}
