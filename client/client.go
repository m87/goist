package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

)


const (
  projectsApi = "/projects"
)


func apiV2(endpoint string) string {
  return "https://api.todoist.com/rest/v2" + endpoint
}

type Project struct {
  Id string `json:"id"`
  Name string `json:"name"`
}

type Task struct {
  Content string `json:"content"`
  ProjectId string `json:"project_id"`
  ProjectName string `json:"project_name"`
}

type Client interface {
  ListProjects() ([]Project, error)
  CreateTask(task Task)
}

func NewTodoistClient(apiKey string) (*TodoistClient, error) {
  var http = &http.Client{}
  var trimmedKey = strings.TrimSpace(apiKey)

  if len(trimmedKey) == 0 {
    return nil, errors.New("Empty api key")
  }

  return &TodoistClient{apiKey: trimmedKey, http: *http}, nil
}

type TodoistClient struct {
  apiKey string
  http http.Client
}

func (t TodoistClient) appendHeaders(req *http.Request) {
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", t.apiKey))
  req.Header.Add("Content-Type", "application/json")
}

func (t TodoistClient) ListProjects() ([]Project, error) {
    req, err := http.NewRequest("GET", apiV2(projectsApi), nil)

    if err != nil {
      log.Fatal("Unable to create projects list request")
      return nil, err
    }

    t.appendHeaders(req)

    resp, err := t.http.Do(req)

    if err != nil {
      log.Fatal("Projects list request failed")
      return nil, err
    }

    body, err := io.ReadAll(resp.Body)
    var projects []Project

    json.Unmarshal(body, &projects)

    log.Print(resp)

    return projects, nil
}

func (t TodoistClient) CreateTask(task Task) {

}
