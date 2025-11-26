package controllers

import "github.com/gin-gonic/gin"

func StartServer() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	//นำตัวcrontroller เข้า Server
	DemoController(router)
	UserController(router)
	PersonControoller(router)
	router.Run()
}
