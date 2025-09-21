package router

import (
	"JHETBackend/internal/common/permission"
	middleware "JHETBackend/internal/middlewares"

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
