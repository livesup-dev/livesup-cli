package models

type User struct {
	ID          string
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string
	AvatarUrl   string `json:"avatar_url"`
	Provider    string
	ConfirmedAt string `json:"confirmed_at"`
	State       string
	InsertedAt  string `json:"inserted_at"`
	UpdatedAt   string `json:"updated_at"`
}

func (user *User) FullName() string {
	return user.FirstName + " " + user.LastName
}

func (user *User) GetID() string {
	return user.ID
}
