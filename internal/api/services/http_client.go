package services

import (
	"encoding/json"
	"fmt"
	"io"
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

// TODO: Im not really proud of making this var public
// this is a way to have dependency injection so I can test
// this package
var (
	Client HttpClient
)

func init() {
	Client = &http.Client{}
}

func doGet(apiResponse ApiResponse, path string) ApiResponse {
	req := newRequest(http.MethodGet, path)

	resp, err := Client.Do(req)

	panicOnError(err)

	buildResponse(resp.Body, &apiResponse)

	return apiResponse
}

func panicOnError(err error) {
	if err != nil {
		panic(fmt.Errorf("fail to perform the request: %w", err))
	}
}

func newRequest(method, path string) *http.Request {
	req, err := http.NewRequest("GET", buildApiPath(path), nil)

	panicOnError(err)

	req.Header.Add("ContentType", contentType)
	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", config.Token()))

	return req
}

func buildResponse(body io.ReadCloser, target *interface{}) *interface{} {
	defer body.Close()

	fullBody, err := ioutil.ReadAll(body)

	panicOnError(err)

	err = json.Unmarshal(fullBody, target)

	panicOnError(err)

	if err != nil { // Parse []byte to go struct pointer
		fmt.Printf("Can not unmarshal JSON: %T\n%s\n%#v\n", err, err, err)
	}

	return target
}

func buildApiPath(path string) string {
	return fmt.Sprintf("%sapi%s", config.URL(), path)
}

func buildApiPathWithId(path string, id string) string {
	return buildApiPath(fmt.Sprintf("%s/%s", path, id))
}
