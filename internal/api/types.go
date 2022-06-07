package api

type User struct {
	ID          string
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string
	AvatarUrl   string `json:"avatar_url"`
	Provider    string
	ConfirmedAt string `json:"confirmed_at"`
	State       string
}
