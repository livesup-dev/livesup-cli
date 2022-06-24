package ui

import "github.com/livesup-dev/livesup-cli/internal/api/models"

type UserTable struct {
	Rows    [][]string
	Headers []string
	Users   []models.User
}

func buildUserTable(users []models.User) UserTable {
	rows := buildUserRows(users)
	return UserTable{
		Users:   users,
		Headers: []string{"#", "First Name", "Last Name", "Email"},
		Rows:    rows,
	}
}

func RenderUserTable(users []models.User) {
	userTable := buildUserTable(users)
	DrawTable(userTable)
}

func (userTable UserTable) GetRows() [][]string {
	return userTable.Rows
}

func (userTable UserTable) GetHeaders() []string {
	return userTable.Headers
}

func buildUserRows(users []models.User) [][]string {
	rows := [][]string{}

	for _, u := range users {
		rows = append(rows, []string{u.ID, u.FirstName, u.LastName, u.Email})
	}

	return rows
}
