package models

type User struct {
	ID          string `json:"id"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	Email       string `json:"email,omitempty"`
	AvatarUrl   string `json:"avatar_url,omitempty"`
	Provider    string `json:"provider,omitempty"`
	ConfirmedAt string `json:"confirmed_at,omitempty"`
	State       string `json:"state,omitempty"`
	InsertedAt  string `json:"inserted_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

func (user *User) FullName() string {
	return user.FirstName + " " + user.LastName
}

func (user *User) GetID() string {
	return user.ID
}
