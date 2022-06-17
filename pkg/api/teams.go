package api

import (
	"context"
	"fmt"

	"github.com/carlmjohnson/requests"
	"github.com/livesup-dev/livesup-cli/pkg/config"
)

type Team struct {
	ID          string
	Name        string
	Slug        string
	Description string
	AvatarUrl   string `json:"avatar_url"`
	InsertedAt  string `json:"inserted_at"`
	UpdatedAt   string `json:"updated_at"`
}

type TeamsResponse struct {
	Teams []Team `json:"data"`
}

// TODO: Is there any way to not repeat
// the body of all these functions and do
// something like apiGet("api/users", &response)?
func GetAllTeams() TeamsResponse {
	var response TeamsResponse
	err := requests.
		URL(config.URL()).
		Pathf("api/teams").
		ContentType("application/json").
		Bearer(config.Token()).
		ToJSON(&response).
		Fetch(context.Background())

	if err != nil {
		panic(fmt.Errorf("fatal error reading API: %w", err))
	}
	return response
}
