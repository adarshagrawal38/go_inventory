package controllers

import (
	"fmt"
	"inventory-management/models"
	"testing"
)

var createItemInDBMock func(item *models.Item)
var findItemInDBMock func(item *models.Item)

type fireQueryMock struct{}

func (f fireQueryMock) createItemInDB(item *models.Item) {
	createItemInDBMock(item)
}
func (f fireQueryMock) findAllItemsInDB(data *models.Items) {
	*data = models.Items{
		{ID: 1},
		{ItemName: "Abcd"},
	}
	fmt.Print("data ", data)
}
func TestCreateItem(t *testing.T) {
	item := models.Item{
		MRP:          130,
		Description:  "abcd",
		ItemName:     "xyz",
		Quantity:     2,
		SellingPrice: 123.9,
	}

	query := fireQueryMock{}
	createItemInDBMock = func(item *models.Item) {
		fmt.Println("Inside mock")
		item.ID = 1
	}

	result := CreateItem(&item, query)
	fmt.Println(result)
	if result.ID != 1 {
		t.Error("Invalid Id")
	}

}

func TestGetAllItems(t *testing.T) {
	item := models.Item{
		MRP:          130,
		Description:  "abcd",
		ItemName:     "xyz",
		Quantity:     2,
		SellingPrice: 123.9,
	}
	fmt.Println(item)

	query := fireQueryMock{}

	result := GetAllItems(query)
	fmt.Println(result)

	if len(result) == 0 {
		t.Error("Unexpected value")

	}

}
