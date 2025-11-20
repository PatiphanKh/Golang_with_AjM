![[Pasted image 20250306150132.png]]

# GIN
- GIN => Web framework for Go
- [https://pkg.go.dev/github.com/gin-gonic/gin](https://pkg.go.dev/github.com/gin-gonic/gin)

```sh
go get github.com/gin-gonic/gin
```

**main.go**
```go
package main

import "github.com/gin-gonic/gin"

func main() {
	// Release mode
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong!!!",
		})
	})
	router.Run()
}

```

![[Pasted image 20250306150529.png]]

```go
c.JSON(200, gin.H{
			"message": "pong!!!",
})

type H map[string]any
H is a shortcut for map[string]any
```

## Test on API

![[Pasted image 20250306150629.png]]

```
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)
```

# Live-reload within the Air

- No white space in application path

```sh
go install github.com/air-verse/air@latest
```

- air => run in live-reload mode

```sh
air
```

![[Pasted image 20250306151610.png]]

# Router
- Separate file

![[Pasted image 20250306151917.png]]

**controller/server.go**
```go
package controller

import "github.com/gin-gonic/gin"

func StartServer() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Api is now working",
		})
	})
	router.Run()
}

```



![[Pasted image 20250306152243.png]]

**controller/demo.go**
```go
package controller

import "github.com/gin-gonic/gin"

func DemoController(router *gin.Engine) {
	router.GET("/ping", ping)
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Pong!!",
	})
}

```


**controller/server.go**
```go
package controller

import "github.com/gin-gonic/gin"

func StartServer() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Api is now working",
		})
	})
	// Include Demo Controller
	DemoController(router)
	router.Run()
}

```

**main.go**
```go
package main

import (
	"go-gin/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	// Release mode
	gin.SetMode(gin.ReleaseMode)
	controller.StartServer()
}

```

![[Pasted image 20250306152900.png]]

