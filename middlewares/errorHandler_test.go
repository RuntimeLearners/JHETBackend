package middleware_test

import (
	"JHETBackend/internal/common/exception"
	middleware "JHETBackend/internal/middlewares"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestUnifiedErrorHandler_Panic(t *testing.T) {

	defer func() {
		r := recover()
		sr, ok := r.(string)
		if !ok || !strings.Contains(sr, "Fail-Fast enabled") {
			t.Errorf("unexpected panic message: %v", r)
		}
		// 如果走到这里，说明 panic 被成功捕获且值正确，测试通过
	}()

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(middleware.UnifiedErrorHandler())
	router.GET("/panic", func(c *gin.Context) {
		panic("test: something went wrong")
	})

	//自发请求，用于测试
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/panic", nil)
	router.ServeHTTP(w, req)
	log.Printf("%v+", w.Body)
}

func TestUnifiedErrorHandler_CErrorExp(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.Use(middleware.UnifiedErrorHandler())
	router.GET("/error", func(c *gin.Context) {
		c.Error(exception.TestIntendedException)
	})

	//自发请求，用于测试
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/error", nil)
	router.ServeHTTP(w, req)
	log.Printf("%v+", w.Body)
}

func TestUnifiedErrorHandler_CErrorUnExp(t *testing.T) {

	// defer func() {
	// 	r := recover()
	// 	sr, ok := r.(string)
	// 	if !ok || !strings.Contains(sr, "Fail-Fast enabled") {
	// 		t.Errorf("unexpected panic message: %v", r)
	// 	}
	// 	// 如果走到这里，说明 panic 被成功捕获且值正确，测试通过
	// }()

	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.Use(middleware.UnifiedErrorHandler())
	router.GET("/error", func(c *gin.Context) {
		c.Error(errors.New("manual error"))
	})

	//自发请求，用于测试
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/error", nil)
	router.ServeHTTP(w, req)
	log.Printf("%v+", w.Body)
}
