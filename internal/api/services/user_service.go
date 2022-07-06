package services

import (
	"github.com/livesup-dev/livesup-cli/internal/api/models"
)

const usersPath = "/users"

type UserService interface {
	All() (*UserList, error)
	Update(user *models.User) (*models.User, error)
	Create(user *models.User) (*models.User, error)
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

func (userSingle *UserSingle) GetModel() models.Model {
	return userSingle.User
}

func (*userService) All() (*UserList, error) {
	usersList, err := doGet(&UserList{}, usersPath)

	if err != nil {
		return nil, err
	}

	return usersList.(*UserList), err
}

func (*userService) Update(user *models.User) (*models.User, error) {
	body := make(map[string]models.Model)
	body["user"] = user

	updatedUser, err := doUpdate(&body, &UserSingle{}, user.GetID(), usersPath)

	if err != nil {
		return nil, err
	}

	return (*updatedUser).(*UserSingle).User, err
}

func (*userService) Create(user *models.User) (*models.User, error) {
	body := make(map[string]models.Model)
	body["user"] = user

	newUser, err := doPost(&body, &UserSingle{}, usersPath)

	if err != nil {
		return nil, err
	}

	return (*newUser).(*UserSingle).User, err
}
