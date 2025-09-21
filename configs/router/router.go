package router

import (
	"JHETBackend/common/permission"
	middleware "JHETBackend/middlewares"

	"github.com/gin-gonic/gin"
)

func InitEngine() *gin.Engine {
	ginEngine := gin.Default()

	ginEngine.GET("/test", middleware.UnifiedErrorHandler(),
		middleware.NeedPerm(
			permission.Perm_ForTestOnly1,
			permission.Perm_ForTestOnly2))

	return ginEngine
}
