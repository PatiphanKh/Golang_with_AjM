package controllers

import (
	"fmt"
	dto "go_gin/DTO"
	"go_gin/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func PersonControoller(router *gin.Engine) {
	routers := router.Group("/person")
	{
		routers.GET("/", getPersons)
		routers.POST("", createPerson)

	}

}

func getPersons(c *gin.Context) {
	person1 := models.Person{
		ID:          1,
		FirstName:   "Potchara",
		LastName:    "P",
		Age:         30,
		Email:       "potchara@example.com",
		Password:    "password123",
		PostAddress: models.Address{HouseNo: "123", City: "Bangkok"},
	}
	person2 := models.Person{
		ID:          2,
		FirstName:   "John",
		LastName:    "Doe",
		Age:         25,
		Email:       "john.doe@example.com",
		Password:    "securepass",
		PostAddress: models.Address{HouseNo: "456", City: "New York"},
	}
	persons := []models.Person{person1, person2}
	c.JSON(200, persons)
	c.JSON(200, gin.H{
		"messagese": "Person Get",
	})
}

func createPerson(c *gin.Context) {
	user := dto.User{}
	err := c.ShouldBindJSON(&user)
	fmt.Printf("%v %v", user.FirstName, user.LastName)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	//c.JSON(200, user)

	modelUser := models.User{}
	copier.Copy(&modelUser, &user)
	c.JSON(200, user)
}
