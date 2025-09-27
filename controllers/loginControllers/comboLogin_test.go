package loginControllers_test

import (
	"JHETBackend/controllers/loginControllers"
	middleware "JHETBackend/middlewares"
	"log"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_ComboLogin(t *testing.T) {

	gin.SetMode(gin.TestMode)
	router := gin.New()

	router.Use(middleware.UnifiedErrorHandler())
	router.POST("/api/auth/login/combo", middleware.UnifiedErrorHandler(), loginControllers.AuthByCombo)

	w := httptest.NewRecorder()
	reqBody := `{"account": "张三", "password": "password123", "userType": "admin", "rememberMe": true}`
	req := httptest.NewRequest("POST", "/api/auth/login/combo",
		strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	log.Printf("%v+", w.Body)
}
