package api

import (
	"context"
	"fmt"

	"github.com/carlmjohnson/requests"
	"github.com/livesup-dev/livesup-cli/pkg/config"
)

type User struct {
	ID          string
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string
	AvatarUrl   string `json:"avatar_url"`
	Provider    string
	ConfirmedAt string `json:"confirmed_at"`
	State       string
	InsertedAt  string `json:"inserted_at"`
	UpdatedAt   string `json:"updated_at"`
}

func (u User) FullName() string {
	return u.FirstName + " " + u.LastName
}

// TODO: Is there any way to not duplicate all
// these types?
type UsersResponse struct {
	Users []User `json:"data"`
}

func GetAllUsers() UsersResponse {
	var response UsersResponse
	err := requests.
		URL(config.URL()).
		Pathf("api/users").
		ContentType("application/json").
		Bearer(config.Token()).
		ToJSON(&response).
		Fetch(context.Background())

	if err != nil {
		panic(fmt.Errorf("fatal error reading API: %w", err))
	}
	return response
}
