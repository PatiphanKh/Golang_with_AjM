package main

import (
	"fmt"
	"go_gorm/models"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	dsn := viper.GetString("mysql.dsn")

	dialector := mysql.Open(dsn)
	db, err := gorm.Open(dialector)
	if err != nil {
		panic("Con Error")
	} else {
		fmt.Println("Con sucess")
	}

	// var products []models.Product
	// db.Preload("OrderItems").Preload("Category").Find(&products)
	// fmt.Println(products)

	// countries := []models.Country{}
	// db.Where("Name like ? or Name = ?", "%ไทย%", "อิตาลี").Find(&countries)
	// fmt.Println(countries)

	// country := models.Country{Name: "บ้านกูเอง" }
	// result := db.Create(&country)
	// if result.Error != nil{
	// 	panic(result.Error)
	// }

	// if result.RowsAffected> 0 {
	// 	fmt.Println("Succes")
	// }
	// countries := []models.Country{}
	// result = db.Find(&countries)
	// if result.Error != nil {
	// 	panic(result.Error)
	// }
	// fmt.Println(countries)

		result := db.Delete(models.Country{}, 366)
	if result.Error != nil {
		panic(result.Error)
	}
	if result.RowsAffected > 0 {
		fmt.Println("Success")
	}

	countries := []models.Country{}
	result = db.Find(&countries)
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Println(countries)
}
