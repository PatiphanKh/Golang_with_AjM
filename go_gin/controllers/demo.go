package controllers

import "github.com/gin-gonic/gin"

func DemoController(router *gin.Engine) {

	routers := router.Group("/ping")
	{
		routers.GET("", ping)
		routers.POST("", POSTping)
		routers.GET("/:name", pingName)
	}

}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"messagese": "Pong controller",
	})
}

func POSTping(c *gin.Context) {
	name := c.PostForm("name")

	c.JSON(200, gin.H{
		"messagese": "Pong controller POST " + name,
	})
}

func pingName(c *gin.Context) {
	name := c.Param("name")
	age := c.Query("age")
	c.JSON(200, gin.H{
		"message": "Pong " + name + " " + age,
	})
}
