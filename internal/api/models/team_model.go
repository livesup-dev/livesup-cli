package models

type Team struct {
	ID          string
	Name        string `json:"name,omitempty"`
	Slug        string `json:"slug,omitempty"`
	Description string `json:"description,omitempty"`
	AvatarUrl   string `json:"avatar_url,omitempty"`
	InsertedAt  string `json:"inserted_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

func (team Team) GetID() string {
	return team.ID
}
