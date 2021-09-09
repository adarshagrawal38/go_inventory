package database

import (
	"fmt"
	"inventory-management/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var DB *gorm.DB
func Connect() {
	DB_USERNAME := "root"
	DB_PASSWORD := "12345"
	DATABASE := "inventory"

	conncetion := fmt.Sprintf("%s:%s@/%s", DB_USERNAME, DB_PASSWORD, DATABASE)

	fmt.Println("Connection string: ", conncetion)
	
	database, err := gorm.Open(mysql.Open(conncetion), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic("Could not connect to db")
	}

	fmt.Println("Connected to databae")

	DB = database
	

	database.AutoMigrate(&models.Users{}, &models.Item{})

}