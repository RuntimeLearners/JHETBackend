package router

import "github.com/gin-gonic/gin"

func InitEngine() *gin.Engine {
	ginEngine := gin.Default()
	return ginEngine
}

func Need(permNames ...string) gin.HandlerFunc {
	var a gin.HandlerFunc
	return a
}
