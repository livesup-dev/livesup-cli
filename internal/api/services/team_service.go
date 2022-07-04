package services

import (
	"context"
	"fmt"

	"github.com/carlmjohnson/requests"
	"github.com/livesup-dev/livesup-cli/internal/api/models"
	"github.com/livesup-dev/livesup-cli/internal/config"
)

const teamsPath = "/teams"

type TeamService interface {
	// Create(team *models.Team) (*models.Team, error)
	All() (*TeamList, error)
	Update(team *models.Team) (*models.Team, error)
}

type teamService struct{}

type TeamList struct {
	Teams []models.Team `json:"data"`
}

func NewTeamService() TeamService {
	return &teamService{}
}

func (*teamService) All() (*TeamList, error) {
	// TODO: errors not implemented
	return doGet(&TeamList{}, teamsPath).(*TeamList), nil
}

type TeamSingle struct {
	Team *models.Team `json:"data"`
}

func (teamSingle *TeamSingle) GetModel() models.Model {
	return teamSingle.Team
}

func (*teamService) Update(team *models.Team) (*models.Team, error) {
	body := make(map[string]models.Model)
	body["team"] = team

	// TODO: How do I actually get rid of all these
	// duplicated code?
	err := requests.
		URL(config.URL()).
		Pathf(buildApiPathWithId(teamsPath, team.GetID())).
		Put().
		BodyJSON(&body).
		ContentType(contentType).
		Bearer(config.Token()).
		ToJSON(&team).
		Fetch(context.Background())

	if err != nil {
		panic(fmt.Errorf("fatal error reading API: %w", err))
	}

	return team, err
}
