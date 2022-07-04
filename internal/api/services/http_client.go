package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/livesup-dev/livesup-cli/internal/api/models"
	"github.com/livesup-dev/livesup-cli/internal/config"
)

const contentType = "application/json"

// Alias
type ApiResponse = interface{}

type Single interface {
	GetModel() models.Model
}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	Client HttpClient
)

func init() {
	Client = &http.Client{}
}

func doGet(apiResponse ApiResponse, path string) ApiResponse {
	req, err := http.NewRequest("GET", buildApiPath(path), nil)

	if err != nil {
		fmt.Printf("Fail to perform the request: %s", err)
		return nil
	}

	req.Header.Add("ContentType", contentType)
	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", config.Token()))
	resp, err := Client.Do(req)
	if err != nil {
		fmt.Printf("Fail to get the data: %s", err)
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("%T\n%s\n%#v\n", err, err, err)
	}

	if err := json.Unmarshal(body, &apiResponse); err != nil { // Parse []byte to go struct pointer
		fmt.Printf("Can not unmarshal JSON: %T\n%s\n%#v\n", err, err, err)
	}

	return apiResponse
}

func buildApiPath(path string) string {
	return fmt.Sprintf("%sapi%s", config.URL(), path)
}

func buildApiPathWithId(path string, id string) string {
	return buildApiPath(fmt.Sprintf("%s/%s", path, id))
}
