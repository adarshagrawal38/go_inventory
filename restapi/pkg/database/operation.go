package database

import (
	"fmt"
	"inventory-management/models"
)

func CreateItemInDB(item *models.Item) {
	fmt.Println("Item created")
	DB.Create(item)
}

func FindAllItems(items *models.Items) {
	DB.Find(items)
}
