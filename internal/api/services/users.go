package services

import "github.com/livesup-dev/livesup-cli/internal/api/models"

const users_path = "/users"

type UserList struct {
	Users []models.User `json:"data"`
}

type UserSingle struct {
	User models.User `json:"data"`
}

func GetAllUsers() *UserList {
	return doGet(&UserList{}, users_path).(*UserList)
}

func UpdateUser(user models.User) models.User {
	return doUpdate(user).GetModel().(models.User)
}
