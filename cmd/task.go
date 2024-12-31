/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/m87/goist/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Parse(input string) (string, string, []string) {
  projectsRegExp := regexp.MustCompile(`#\w+?($|\s+)`)
  labelsRegExp := regexp.MustCompile(`@\w+?($|\s+)`)
  project := strings.TrimPrefix(strings.TrimSpace(projectsRegExp.FindString(input)), "#")
  labels := labelsRegExp.FindAllString(input, 1000)

  // TODO remove @
  
  content := strings.TrimSpace(projectsRegExp.ReplaceAllString(input, ""))
  content = strings.TrimSpace(labelsRegExp.ReplaceAllString(content, ""))
  
  return content, project, labels
}

type Task struct {
  Content string `json:"content"`
  ProjectId string `json:"project_id"`
  Labels []string `json:"labels"`
}


// taskCmd represents the task command
var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Add new task",
	Long: `Add new task. E.g. 
  goist create task "new task #porject @label"
  goist create task "new task" -p project -l label`,
	Run: func(cmd *cobra.Command, args []string) {
    client2 := &http.Client{}

    content, pname, labels := Parse(args[0])

    var project client.Project
    cli, _ := cmd.Context().Value(Client).(client.Client)

    projects, err := cli.ListProjects()

    if err != nil {
      log.Fatal("Unable to list projects") 
    }
    

    for _, p := range projects {
      if p.Name == pname {
        project = p;
        break
      }

    }

    j := &Task{Content: content, ProjectId: project.Id, Labels: labels}

    payload, err := json.Marshal(j)
    log.Print(string(payload))
    if err != nil {
      log.Fatal("Unable to prepare request")
    }

    req, err := http.NewRequest("POST", "https://api.todoist.com/rest/v2/tasks", bytes.NewBuffer(payload))

    log.Print(req)

    if err != nil {
      log.Fatal("Unable to create new task request")
    }

    req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", viper.Get("token")))
    req.Header.Add("Content-Type", "application/json")

    resp, err := client2.Do(req)
    
    if resp != nil {
      log.Fatal(resp)
    }

    if err != nil {
      log.Fatal("Unable to create new task")
    }
		fmt.Println("task called")
	},
}

func init() {
	createCmd.AddCommand(taskCmd)
  taskCmd.Flags().String("content", "c", "New task")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// taskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// taskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
