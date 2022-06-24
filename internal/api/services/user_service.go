package services

import "github.com/livesup-dev/livesup-cli/internal/api/models"

const users_path = "/users"

type UserService interface {
	All() (*UserList, error)
	Update(user *models.User) (*models.User, error)
}

type userService struct{}

func NewUserService() UserService {
	return &userService{}
}

type UserList struct {
	Users []models.User `json:"data"`
}

type UserSingle struct {
	User *models.User `json:"data"`
}

func (*userService) All() (*UserList, error) {
	return doGet(&UserList{}, users_path).(*UserList), nil
}

func (*userService) Update(user *models.User) (*models.User, error) {
	return doUpdate(user, users_path).GetModel().(*models.User), nil
}
