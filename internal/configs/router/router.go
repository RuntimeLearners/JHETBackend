package router

import "github.com/gin-gonic/gin"

func InitEngine() *gin.Engine {
	ginEngine := gin.Default()
	return ginEngine
}
