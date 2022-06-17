package ui

import (
	"github.com/livesup-dev/livesup-cli/pkg/api"
)

type TeamTable struct {
	Rows    [][]string
	Headers []string
	Teams   []api.Team
}

func BuildTeamTable(teams []api.Team) TeamTable {
	rows := buildTeamRows(teams)
	return TeamTable{
		Teams:   teams,
		Headers: []string{"#", "Name", "Inserted At"},
		Rows:    rows,
	}
}

func (teamTable TeamTable) GetRows() [][]string {
	return teamTable.Rows
}

func (teamTable TeamTable) GetHeaders() []string {
	return teamTable.Headers
}

func buildTeamRows(teams []api.Team) [][]string {
	rows := [][]string{}

	for _, t := range teams {
		rows = append(rows, []string{t.ID, t.Name, t.InsertedAt})
	}

	return rows
}
