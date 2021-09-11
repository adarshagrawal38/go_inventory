package database

import (
	"fmt"
	"inventory-management/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// DB_USERNAME := "root"
	// DB_PASSWORD := "12345"
	// DATABASE := "inventory"
	// dsn := "root:12345@tcp(127.0.0.1:3306)/training?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:12345@tcp(inventory-management_database_1)/inventory?charset=utf8mb4&parseTime=True&loc=Local"
	// conncetion := fmt.Sprintf("%s:%s@/%s", DB_USERNAME, DB_PASSWORD, DATABASE)

	fmt.Println("Connection string: ", dsn)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic("Could not connect to db")
	}

	fmt.Println("Connected to databae")

	DB = database

	database.AutoMigrate(&models.Users{}, &models.Item{})

}
