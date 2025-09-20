package main

import "JHETBackend/internal/configs/router"

//import "JHETBackend/internal/configs/database"

func main() {
	ginEng := router.InitEngine()
	ginEng.Run(":8080")
}
