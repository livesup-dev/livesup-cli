package api

type Team struct {
	ID          string
	Name        string
	Slug        string
	Description string
	AvatarUrl   string `json:"avatar_url"`
	Labels      []string
}
