package main

import (
	"JHETBackend/internal/configs/router"

	"github.com/gin-gonic/gin"
)

//import "JHETBackend/internal/configs/database"

func main() {
	ginEng := router.InitEngine()
	ginEng.Run(":8080")
	ginEng.GET("/user", userHandler)

}

func userHandler(c *gin.Context) {
}
