package ui

import "github.com/livesup-dev/livesup-cli/internal/api/models"

type TeamTable struct {
	Rows    [][]string
	Headers []string
	Teams   []models.Team
}

func buildTeamTable(teams []models.Team) TeamTable {
	rows := buildTeamRows(teams)
	return TeamTable{
		Teams:   teams,
		Headers: []string{"#", "Name", "Slug", "Description", "Updated At"},
		Rows:    rows,
	}
}

func RenderTeamTable(teams []models.Team) {
	teamTable := buildTeamTable(teams)
	DrawTable(teamTable)
}

func (teamTable TeamTable) GetRows() [][]string {
	return teamTable.Rows
}

func (teamTable TeamTable) GetHeaders() []string {
	return teamTable.Headers
}

func buildTeamRows(teams []models.Team) [][]string {
	rows := [][]string{}

	for _, t := range teams {
		rows = append(rows, []string{t.ID, t.Name, t.Slug, t.Description, t.UpdatedAt})
	}

	return rows
}
