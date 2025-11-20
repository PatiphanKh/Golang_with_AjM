
- Model => Represent Data Table from database
- DTO (Data Transfer Object) => Represent Data in/out API


# Model

![[Pasted image 20250306162556.png]]

**model/person.go**
```go
package model

type Person struct {
	ID          int
	FirstName   string
	LastName    string
	Age         int
	Email       string
	Password    string
	PostAddress Address
}

type Address struct {
	HouseNo string
	City    string
}

```

# Sample API

**controller/person.go**
```go
package controller

import (
	"go-gin/model"

	"github.com/gin-gonic/gin"
)

func PersonController(router *gin.Engine) {
	routes := router.Group("/person")
	{
		routes.GET("", getAllPersons)
	}
}

func getAllPersons(c *gin.Context) {
	person1 := model.Person{
		ID:          1,
		FirstName:   "Potchara",
		LastName:    "P",
		Age:         30,
		Email:       "potchara@example.com",
		Password:    "password123",
		PostAddress: model.Address{HouseNo: "123", City: "Bangkok"},
	}
	person2 := model.Person{
		ID:          2,
		FirstName:   "John",
		LastName:    "Doe",
		Age:         25,
		Email:       "john.doe@example.com",
		Password:    "securepass",
		PostAddress: model.Address{HouseNo: "456", City: "New York"},
	}
	persons := []model.Person{person1, person2}
	c.JSON(200, persons)
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
	PersonController(router)
	router.Run()
}

```

![[Pasted image 20250306164256.png]]

**Password is also shown**
- Password MUST Not be transfer

# Ignore fields with JSON tag
- `json:"-"` => hide field
- `json:"Field Name"` => specific field name

**model/person.go**
```go
package model

type Person struct {
	ID          int
	FirstName   string
	LastName    string
	Age         int `json:"-"`
	Email       string
	Password    string  `json:"-"`
	PostAddress Address `json:"Address"`
}

type Address struct {
	HouseNo string
	City    string
}
```

![[Pasted image 20250306165652.png]]

# DTO
- JSON Data In and Out 
- Data mapping between JSON and Go 
- Specific Data fields
- Separate from Model

![[Pasted image 20250306211012.png]]

**dto/user.go**
```go
package dto

type User struct {
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	Email     string `json:"email"`
}
```


**controller/person.go**
```go
func createPerson(c *gin.Context) {
	user := dto.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}
```

![[Pasted image 20250306211226.png]]
## Go (Model) to Json
- Transfer DTO to a Model
- Auto Copy
- https://pkg.go.dev/github.com/jinzhu/copier

```sh
go get github.com/jinzhu/copier
```

- `copier.Copy(&to, &from)`

**controller/person.go**
```go
func getAllPersons(c *gin.Context) {
	person1 := model.Person{
		ID:          1,
		FirstName:   "Potchara",
		LastName:    "P",
		Age:         30,
		Email:       "potchara@example.com",
		Password:    "password123",
		PostAddress: model.Address{HouseNo: "123", City: "Bangkok"},
	}
	person2 := model.Person{
		ID:          2,
		FirstName:   "John",
		LastName:    "Doe",
		Age:         25,
		Email:       "john.doe@example.com",
		Password:    "securepass",
		PostAddress: model.Address{HouseNo: "456", City: "New York"},
	}
	persons := []model.Person{person1, person2}
	users := []dto.User{}
	copier.Copy(&users, &persons)
	// personDto := dto.PersonDTO{}

	c.JSON(200, users)

}
```

![[Pasted image 20250306212556.png]]


# DTO to Model

- Receive json with fields (fname, lname, email)
- Map to struct fields FirstName, LastName, Email

## JSON To Go (Model)
**controller/person.go**
```go
func createPerson(c *gin.Context) {
	user := dto.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// c.JSON(200, user)
	person := model.Person{}
	copier.Copy(&person, &user)
	c.JSON(200, person)
}
```

![[Pasted image 20250306211955.png]]

