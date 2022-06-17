package ui

import (
	"github.com/livesup-dev/livesup-cli/pkg/api"
)

type UserTable struct {
	Rows    [][]string
	Headers []string
	Users   []api.User
}

func BuildUserTable(users []api.User) UserTable {
	rows := buildUserRows(users)
	return UserTable{
		Users:   users,
		Headers: []string{"#", "First Name", "Last Name", "Email"},
		Rows:    rows,
	}
}

func (userTable UserTable) GetRows() [][]string {
	return userTable.Rows
}

func (userTable UserTable) GetHeaders() []string {
	return userTable.Headers
}

func buildUserRows(users []api.User) [][]string {
	rows := [][]string{}

	for _, u := range users {
		rows = append(rows, []string{u.ID, u.FirstName, u.LastName, u.Email})
	}

	return rows
}
