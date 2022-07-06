package services_test

//https://stackoverflow.com/questions/19998250/proper-package-naming-for-testing-with-the-go-language
import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/livesup-dev/livesup-cli/internal/api/models"
	"github.com/livesup-dev/livesup-cli/internal/api/services"
	"github.com/livesup-dev/livesup-cli/internal/utils/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNewTeamService(t *testing.T) {
	service := services.NewTeamService()
	assert.NotNil(t, service)
}

func TestTeamService(t *testing.T) {
	t.Run("all with error", func(t *testing.T) {
		mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
			return nil, errors.New(
				"Error from web server",
			)
		}

		services.Client = &mocks.MockClient{}

		service := services.NewTeamService()

		assert.Panics(t, func() {
			data, _ := service.All()
			assert.Nil(t, data)
		}, "Error from web server")
	})

	t.Run("all", func(t *testing.T) {
		json := `{"data":[{"avatar_url":"https://pythiabot.s3.amazonaws.com/teams/customer-portal.png","description":"this guat?","id":"d61f5ae8-5cf3-4290-9c4a-dae8ed91eb60","inserted_at":"2022-06-15T10:46:49","name":"aaa","slug":"customer-portal","updated_at":"2022-06-24T13:57:23"},{"avatar_url":"https://pythiabot.s3.amazonaws.com/teams/customer-success.png","description":null,"id":"9c6dc806-9d50-4496-a393-b3fc8f967b91","inserted_at":"2022-06-15T10:46:49","name":"Customer Success","slug":"customer-success","updated_at":"2022-06-15T10:46:49"}]}`
		r := ioutil.NopCloser(bytes.NewReader([]byte(json)))

		mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		}

		services.Client = &mocks.MockClient{}

		service := services.NewTeamService()
		teamList, err := service.All()

		assert.Nil(t, err)
		assert.Equal(t, len(teamList.Teams), 2)
	})

	t.Run("update", func(t *testing.T) {
		json := `{"data":{"avatar_url":"https://pythiabot.s3.amazonaws.com/teams/customer-portal.png","description":"New desc","id":"d61f5ae8-5cf3-4290-9c4a-dae8ed91eb60","inserted_at":"2022-06-15T10:46:49","name":"aaa","slug":"customer-portal","updated_at":"2022-07-04T23:19:33"}}`
		r := ioutil.NopCloser(bytes.NewReader([]byte(json)))

		mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		}

		services.Client = &mocks.MockClient{}

		service := services.NewTeamService()

		team := &models.Team{
			ID:          "d61f5ae8-5cf3-4290-9c4a-dae8ed91eb60",
			Description: "New desc",
		}
		updatedTeam, err := service.Update(team)

		assert.Nil(t, err)
		assert.Equal(t, updatedTeam.Description, "New desc")
	})
}
