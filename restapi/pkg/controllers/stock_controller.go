package controllers

import (
	"inventory-management/models"
	"inventory-management/restapi/operations/stock"
	"inventory-management/restapi/pkg/database"
)

func CreateItem(params stock.PostInventoryParams) *models.Item {
	var item *models.Item = params.Body
	database.DB.Create(item)

	return item
}

func GetAllItems() models.Items {
	items := models.Items{}
	database.DB.Find(&items)

	return items
}

func UdateAnItem(params stock.PatchInventoryParams) *models.Item {
	var item *models.Item = params.Body
	database.DB.Updates(item)

	return item
}

func DeleteAnItem(params stock.DeleteInventoryItemIDParams) error {
	item := models.Item{
		ID: params.ItemID,
	}
	err := database.DB.Delete(&item)

	return err.Error
}
