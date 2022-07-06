package services

import (
	"github.com/livesup-dev/livesup-cli/internal/api/models"
)

const teamsPath = "/teams"

type TeamService interface {
	// Create(team *models.Team) (*models.Team, error)
	All() (*TeamList, error)
	Update(team *models.Team) (*models.Team, error)
	Create(team *models.Team) (*models.Team, error)
}

type teamService struct{}

type TeamList struct {
	Teams []models.Team `json:"data"`
}

func NewTeamService() TeamService {
	return &teamService{}
}

func (*teamService) All() (*TeamList, error) {
	teamList, err := doGet(&TeamList{}, teamsPath)

	if err != nil {
		return nil, err
	}
	return teamList.(*TeamList), err
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

	doUpdate(&body, &TeamSingle{}, team.GetID(), teamsPath)

	return team, nil
}

func (*teamService) Create(team *models.Team) (*models.Team, error) {
	body := make(map[string]models.Model)
	body["team"] = team

	newTeam, err := doPost(&body, &TeamSingle{}, teamsPath)

	if err != nil {
		return nil, err
	}

	return (*newTeam).(*TeamSingle).Team, err
}
