package services_test

//https://stackoverflow.com/questions/19998250/proper-package-naming-for-testing-with-the-go-language
import (
	"testing"

	"github.com/livesup-dev/livesup-cli/internal/api/services"
	"github.com/stretchr/testify/assert"
)

func TestNewTeamService(t *testing.T) {
	service := services.NewTeamService()
	assert.NotNil(t, service)
}

func TestTeamService(t *testing.T) {
	t.Run("all", func(t *testing.T) {
		assert.Equal(t, 1, 1)
	})
}
