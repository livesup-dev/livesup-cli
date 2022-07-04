package services

import (
	"context"
	"fmt"

	"github.com/carlmjohnson/requests"
	"github.com/livesup-dev/livesup-cli/internal/api/models"
	"github.com/livesup-dev/livesup-cli/internal/config"
)

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

func (userSingle *UserSingle) GetModel() models.Model {
	return userSingle.User
}

func (*userService) All() (*UserList, error) {
	return doGet(&UserList{}, users_path).(*UserList), nil
}

func (*userService) Update(user *models.User) (*models.User, error) {
	body := make(map[string]models.Model)
	body["user"] = user

	// TODO: How do I actually get rid of all these
	// duplicated code?
	err := requests.
		URL(config.URL()).
		Pathf(buildApiPathWithId(teamsPath, user.GetID())).
		Put().
		BodyJSON(&body).
		ContentType(contentType).
		Bearer(config.Token()).
		ToJSON(&user).
		Fetch(context.Background())

	if err != nil {
		panic(fmt.Errorf("fatal error reading API: %w", err))
	}

	return user, err
}
