/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)


type Project struct {
  Name string 
  Id string
}


func ListProjects() []Project {
client := &http.Client{}
    req, err := http.NewRequest("GET", "https://api.todoist.com/rest/v2/projects", nil)

    if err != nil {
      log.Fatal("Unable to create projects list request")
    }
    req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", viper.Get("token")))
    req.Header.Add("Content-Type", "application/json")
    resp, err := client.Do(req)

    if err != nil {
      log.Fatal("Projects list request failed")
    }

    body, err := io.ReadAll(resp.Body)
    var projects []Project

    json.Unmarshal(body, &projects)

    return projects

}


// projectCmd represents the project command
var projectListCmd = &cobra.Command{
	Use:   "project",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("project called")
    projects := ListProjects()
    for _, project := range projects {
      fmt.Printf("%s(%s)\n", project.Name, project.Id)
    }

	},
}

func init() {
	listCmd.AddCommand(projectListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// projectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// projectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
