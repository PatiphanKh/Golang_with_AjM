

![[Pasted image 20250306153343.png]]

## Router/Controller
1. Http method
2. Path/Endpoint
3. Handler function
	1. Function Signature
4. Return value
	1. String/JSON/Others

**controller/demo.go**
```go
package controller

import "github.com/gin-gonic/gin"

func DemoController(router *gin.Engine) {
	router.GET("/ping", ping)
	router.POST("/ping", pingPost)
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Pong!!",
	})
}

func pingPost(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "POST Pong!!",
	})
}

```

![[Pasted image 20250306155542.png]]

# Grouped Router

**controller/demo.go**
```go
func DemoController(router *gin.Engine) {
	routes := router.Group("/ping")
	{
		routes.GET("", ping)
		routes.POST("", pingPost)
	}
}
```

# Receiving Data

## Path Parameters
- `name := c.Param("name")`

**controller/demo.go**
```go
package controller

import "github.com/gin-gonic/gin"

func DemoController(router *gin.Engine) {
	routes := router.Group("/ping")
	{
		routes.GET("/", ping)
		routes.POST("/", pingPost)
		routes.GET("/:name", pingName)
	}
}

func pingName(c *gin.Context) {
	name := c.Param("name")
	c.JSON(200, gin.H{
		"message": "Pong " + name,
	})
}
```

![[Pasted image 20250306160038.png]]

## Query Parameter
- ?

```go
func pingName(c *gin.Context) {
	name := c.Param("name")
	age := c.Query("age")
	c.JSON(200, gin.H{
		"message": "Pong " + name + " " + age,
	})
}
```


![[Pasted image 20250306160254.png]]

## PostForm Parameter
- HTML Form variable
- `name := c.PostForm("name")`
- `age := c.DefaultPostForm("age", "18")`

```go
func pingPost(c *gin.Context) {
	name := c.PostForm("name")
	age := c.DefaultPostForm("age", "18")
	c.JSON(200, gin.H{
		"message": "POST Pong!! " + name + " " + age,
	})
}
```

![[Pasted image 20250306160843.png]]


