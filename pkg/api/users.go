package api

import (
	"context"
	"fmt"

	"github.com/carlmjohnson/requests"
	"github.com/livesup-dev/livesup-cli/pkg/config"
)

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
