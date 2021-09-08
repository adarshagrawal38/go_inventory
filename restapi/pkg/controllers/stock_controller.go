package controllers

import (
	"fmt"
	"inventory-management/models"
	"inventory-management/restapi/operations/stock"
	"inventory-management/restapi/pkg/database"
)


func CreateItem(params stock.PostInventoryParams) models.Item {
	item := models.Item {
		MRP: params.Price,
		Description: "item",
		ItemName: params.ItemName,
		Quantity: 1,
		SellingPrice: params.Price,
	}
	database.DB.Create(&item)
	fmt.Println(item)
	return item

}