package services

import (
	"context"
	"fmt"

	"github.com/carlmjohnson/requests"
	"github.com/livesup-dev/livesup-cli/internal/api/models"
	"github.com/livesup-dev/livesup-cli/internal/config"
)

const contentType = "application/json"

// Alias
type ApiResponse = interface{}

type Single interface {
	GetModel() *models.Model
}

func doGet(apiResponse ApiResponse, path string) ApiResponse {
	err := requests.
		URL(config.URL()).
		Pathf(buildApiPath(path)).
		ContentType(contentType).
		Bearer(config.Token()).
		ToJSON(&apiResponse).
		Fetch(context.Background())

	if err != nil {
		panic(fmt.Errorf("fatal error reading API: %w", err))
	}
	return apiResponse
}

func doUpdate(model models.Model, path string) Single {
	// TODO: Im not proud of these lines
	var modelName string
	var singleInterface interface{}

	switch model.(type) {
	case *models.Team:
		modelName = "team"
		singleInterface = &TeamSingle{}
	case *models.User:
		modelName = "user"
		singleInterface = &UserSingle{}
	}
	body := make(map[string]models.Model)
	body[modelName] = model

	err := requests.
		URL(config.URL()).
		Pathf(buildApiPathWithId(path, model.GetID())).
		Put().
		BodyJSON(&body).
		ContentType(contentType).
		Bearer(config.Token()).
		ToJSON(&singleInterface).
		Fetch(context.Background())

	if err != nil {
		panic(fmt.Errorf("fatal error reading API: %w", err))
	}
	return singleInterface
}

func buildApiPath(path string) string {
	return "api/" + path
}

func buildApiPathWithId(path string, id string) string {
	return buildApiPath(fmt.Sprintf("%s/%s", path, id))
}
