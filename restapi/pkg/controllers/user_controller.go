package controllers

import (
	"fmt"
	"inventory-management/models"
	"inventory-management/restapi/operations/user"
	"inventory-management/restapi/pkg/database"
)


func CreateUser(params user.PostUserParams) {

	var users *models.Users = params.Body
	database.DB.Create(users)
	fmt.Println(users)
}