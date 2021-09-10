package controllers

import (
	"fmt"
	"inventory-management/models"
	"inventory-management/restapi/operations/authorization"
	"inventory-management/restapi/operations/user"
	"inventory-management/restapi/pkg/database"
)

func CreateUser(params user.PostUserParams) *models.Users {

	var users *models.Users = params.Body
	database.DB.Create(users)
	fmt.Println(users)
	return users
}

func DeleteUser(params user.DeleteUserIDParams) error {
	user := models.Users{
		ID: params.ID,
	}
	err := database.DB.Delete(&user)
	fmt.Println(user)
	fmt.Println(err.Error)
	return err.Error
}

func FindUserById(params user.GetUserIDParams) models.Users {
	user := models.Users{
		ID: params.ID,
	}
	database.DB.Find(&user)
	return user
}

func UpdateUser(params user.PutUserParams) *models.Users {
	var users *models.Users = params.Body
	database.DB.Updates(users)
	fmt.Println(users)
	return users
}

func LogInUser(params authorization.PostLoginParams) int {
	
	count := int64(0)
	user := models.Users{}
	email1 := params.Body.Email
	pass1 := params.Body.Password
	database.DB.Model(user).Where("email = ? AND `password` = ?", email1, pass1).Count(&count)
	if count > 0 {
		return 1
	} else {
		return 0
	}
}
