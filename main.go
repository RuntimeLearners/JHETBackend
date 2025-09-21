package main

import (
	"JHETBackend/configs/router"

	"github.com/gin-gonic/gin"
)

//import "JHETBackend/configs/database"

func main() {
	ginEng := router.InitEngine()
	ginEng.Run(":8080")
	ginEng.GET("/user", userHandler)

}

func userHandler(c *gin.Context) {
}
