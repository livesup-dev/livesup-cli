package api

import (
	"context"
	"fmt"

	"github.com/carlmjohnson/requests"
	"github.com/livesup-dev/livesup-cli/internal/config"
)

type Team struct {
	ID          string
	Name        string `json:"name,omitempty"`
	Slug        string `json:"slug,omitempty"`
	Description string `json:"description,omitempty"`
	AvatarUrl   string `json:"avatar_url,omitempty"`
	InsertedAt  string `json:"inserted_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

type TeamsResponse struct {
	Teams []Team `json:"data"`
}

type ApiResponse = interface{}

// TODO: Is there any way to not repeat
// the body of all these functions and do
// something like apiGet("api/users", &response)?
func GetAllTeams() TeamsResponse {
	var response TeamsResponse // = &TeamsResponse{}
	// data, ok = getResponse(response).(*TeamsResponse)

	return getResponse(response).(TeamsResponse)
}

func getResponse(apiResponse ApiResponse) ApiResponse {
	err := requests.
		URL(config.URL()).
		Pathf("api/teams").
		ContentType("application/json").
		Bearer(config.Token()).
		ToJSON(&apiResponse).
		Fetch(context.Background())

	if err != nil {
		panic(fmt.Errorf("fatal error reading API: %w", err))
	}
	return apiResponse
}

func UpdateTeam(team Team) Team {
	teamBody := make(map[string]Team)
	teamBody["team"] = team

	err := requests.
		URL(config.URL()).
		Pathf("api/teams/%s", team.ID).
		Put().
		BodyJSON(&teamBody).
		ContentType("application/json").
		Bearer(config.Token()).
		ToJSON(&team).
		Fetch(context.Background())

	if err != nil {
		panic(fmt.Errorf("fatal error reading API: %w", err))
	}
	return team
}
