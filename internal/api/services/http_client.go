package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

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
	Client = &http.Client{
		Timeout: 5 * time.Second,
	}
}

func doGet(apiResponse ApiResponse, path string) (ApiResponse, error) {
	req := newRequest(http.MethodGet, path)

	resp, err := Client.Do(req)

	panicOnError(err)

	if resp.StatusCode == http.StatusInternalServerError {
		return nil, errors.New("internal server error")
	}

	if resp.StatusCode != http.StatusOK {
		b, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("fail to perform the request: %s", string(b))
	}

	buildResponse(resp.Body, &apiResponse)

	return apiResponse, nil
}

// TODO: can't we use generics to deal with the "target" param?
func doUpdate(body *map[string]models.Model, target interface{}, id string, path string) (*interface{}, error) {
	path = buildApiPathWithId(path, id)
	req := newRequestWithBody(http.MethodPut, buildApiPath(path), body)

	resp, err := Client.Do(req)

	panicOnError(err)

	// fmt.Println(resp.StatusCode)
	// b, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(b))

	// TODO: Create a function to handle the http respose
	if resp.StatusCode == http.StatusInternalServerError {
		return nil, errors.New("internal server error")
	}
	if resp.StatusCode != http.StatusOK {
		b, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("fail to perform the request: %s", string(b))
	}

	buildResponse(resp.Body, &target)

	return &target, nil
}

func doPost(body *map[string]models.Model, target interface{}, path string) (*interface{}, error) {
	req := newRequestWithBody(http.MethodPost, buildApiPath(path), body)

	resp, err := Client.Do(req)

	panicOnError(err)

	if resp.StatusCode == http.StatusInternalServerError {
		return nil, errors.New("internal server error")
	}

	if resp.StatusCode != http.StatusCreated {
		b, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("fail to perform the request: %s", string(b))
	}

	buildResponse(resp.Body, &target)

	return &target, nil
}

func panicOnError(err error) {
	if err != nil {
		panic(fmt.Errorf("fail to perform the request: %w", err))
	}
}

func newRequest(method, path string) *http.Request {
	req, err := http.NewRequest(method, buildApiPath(path), nil)

	panicOnError(err)

	addHeaders(req)

	return req
}

func newRequestWithBody(method, path string, body interface{}) *http.Request {
	jsonBytes, err := json.Marshal(body)
	// fmt.Println(string(jsonBytes))
	panicOnError(err)

	req, err := http.NewRequest(method, path, bytes.NewBuffer(jsonBytes))

	addHeaders(req)

	panicOnError(err)

	return req
}

func addHeaders(req *http.Request) {
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", config.Token()))
}

func buildResponse(body io.ReadCloser, target *interface{}) *interface{} {
	defer body.Close()

	fullBody, err := ioutil.ReadAll(body)
	// fmt.Println(fullBody)
	panicOnError(err)

	err = json.Unmarshal(fullBody, target)

	if err != nil { // Parse []byte to go struct pointer
		fmt.Printf("Can not unmarshal JSON: %T\n%s\n%#v\n", err, err, err)
	}

	return target
}

func buildApiPath(path string) string {
	return fmt.Sprintf("%sapi%s", config.URL(), path)
}

func buildApiPathWithId(path string, id string) string {
	return fmt.Sprintf("%s/%s", path, id)
}
