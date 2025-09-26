package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// JsonResponse 返回json格式数据
func JsonResponse(c *gin.Context, httpStatusCode int, code int, msg string, data any) {
	c.JSON(httpStatusCode, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

// 返回成功json格式数据
func JsonSuccessResponse(c *gin.Context, message string, data any) {
	JsonResponse(c, http.StatusOK, 200, message, data)
}

// 返回错误json格式数据
func JsonErrorResponse(c *gin.Context, code int, message string) {
	JsonResponse(c, http.StatusOK, code, message, nil)
}
