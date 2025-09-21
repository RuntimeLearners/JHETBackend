package router

import (
	"JHETBackend/common/permission"
	middleware "JHETBackend/middlewares"

	"github.com/gin-gonic/gin"
)

func SayHello(c *gin.Context) {
	// 200 表示 HTTP 响应状态码（<=> http.StatusOK）
	// 使用 Context 的 String 函数将 "Hello 精弘!" 这句话以纯文本（字符串）的形式返回给前端
	// 实际上是对返回响应的封装
	c.String(200, "Hello go!")
}

func InitEngine() *gin.Engine {
	ginEngine := gin.Default()

	ginEngine.GET("/test", middleware.UnifiedErrorHandler(),
		middleware.NeedPerm(
			permission.Perm_ForTestOnly1,
			permission.Perm_ForTestOnly2), SayHello)

	ginEngine.GET("/api/auth/login/combo", middleware.UnifiedErrorHandler(), SayHello)

	ginEngine.GET("/api/auth/register", middleware.UnifiedErrorHandler(), SayHello)

	return ginEngine
}
