package services

import "github.com/livesup-dev/livesup-cli/internal/api/models"

type Service interface {
	Create(team *models.Model) (*models.Model, error)
	All() (*[]models.Model, error)
}
