package controllers

import (
	"fmt"
	"inventory-management/models"
	"inventory-management/restapi/operations/stock"
	"inventory-management/restapi/pkg/database"
)

type DatabaseInteraction interface {
	createItemInDB(*models.Item)
	findAllItemsInDB(*models.Items)
}
type fireQuery struct{}

func (f fireQuery) createItemInDB(item *models.Item) {
	database.CreateItemInDB(item)
}
func (f fireQuery) findAllItemsInDB(items *models.Items) {
	database.FindAllItems(items)
}
func NewFireQuery() DatabaseInteraction {
	return fireQuery{}
}
func CreateItem(item *models.Item, query DatabaseInteraction) *models.Item {
	query.createItemInDB(item)
	return item
}

func GetAllItems(query DatabaseInteraction) models.Items {
	items := models.Items{}
	query.findAllItemsInDB(&items)
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

func SearchItem(params stock.GetInventorySearchItemNameParams) (*models.Items, error) {
	var items models.Items
	err := database.DB.Where("item_name LIKE ?", "%"+params.ItemName+"%").Find(&items)
	fmt.Println(items)
	return &items, err.Error

}
