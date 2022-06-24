package services

import (
	"github.com/livesup-dev/livesup-cli/internal/api/models"
)

const teams_path = "/teams"

type TeamList struct {
	Teams []models.Team `json:"data"`
}

type TeamSingle struct {
	Team models.Team `json:"data"`
}

func (teamSingle TeamSingle) GetModel() models.Model {
	return teamSingle.Team
}

func GetAllTeams() *TeamList {
	return doGet(&TeamList{}, teams_path).(*TeamList)
}

func UpdateTeam(team models.Team) models.Team {
	// TODO: missing error handler
	return doUpdate(team).GetModel().(models.Team)
}
