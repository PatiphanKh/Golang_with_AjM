package controllers

import "github.com/gin-gonic/gin"

func UserController(router *gin.Engine) {
	router.GET("/user", getuser)
}

func getuser(c *gin.Context) {
	c.JSON(200, gin.H{
		"text": "User Name",
	})
}

