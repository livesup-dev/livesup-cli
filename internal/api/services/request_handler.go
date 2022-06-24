package services

import (
	"context"
	"fmt"

	"github.com/carlmjohnson/requests"
	"github.com/livesup-dev/livesup-cli/internal/api/models"
	"github.com/livesup-dev/livesup-cli/internal/config"
)

const content_type = "application/json"

// Alias
type ApiResponse = interface{}

type Single interface {
	GetModel() models.Model
}

func doGet(apiResponse ApiResponse, path string) ApiResponse {
	err := requests.
		URL(config.URL()).
		Pathf(buildApiPath(path)).
		ContentType(content_type).
		Bearer(config.Token()).
		ToJSON(&apiResponse).
		Fetch(context.Background())

	if err != nil {
		panic(fmt.Errorf("fatal error reading API: %w", err))
	}
	return apiResponse
}

func doUpdate(model models.Model) Single {
	teamBody := make(map[string]models.Model)
	teamBody["team"] = model
	teamSingle := TeamSingle{}

	err := requests.
		URL(config.URL()).
		Pathf("api/teams/%s", model.GetID()).
		Put().
		BodyJSON(&teamBody).
		ContentType("application/json").
		Bearer(config.Token()).
		ToJSON(&teamSingle).
		Fetch(context.Background())

	if err != nil {
		panic(fmt.Errorf("fatal error reading API: %w", err))
	}
	return teamSingle
}

func buildApiPath(path string) string {
	return "api/" + path
}
