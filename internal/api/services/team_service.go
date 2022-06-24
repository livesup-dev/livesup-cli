package services

import (
	"github.com/livesup-dev/livesup-cli/internal/api/models"
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
	// TODO: missing error handler
	return doUpdate(team, teamsPath).GetModel().(*models.Team), nil
}
