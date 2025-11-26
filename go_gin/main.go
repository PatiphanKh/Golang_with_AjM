package main

import (
	"go_gin/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	// //ProductionMode
	gin.SetMode(gin.ReleaseMode)

	// //เส้นทาง
	router := gin.Default()

	// //จัดการเส้นทาง
	// router.GET("/ping", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"messagedr": "Pong",
	// 	})
	// })

	controllers.StartServer()
	router.Run()
}
