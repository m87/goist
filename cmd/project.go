/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/m87/goist/client"
	"github.com/spf13/cobra"
)

// projectCmd represents the project command
var projectListCmd = &cobra.Command{
	Use:   "project",
	Short: "List projects",
	Long: `List projects`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("project called")

    cli, _ := cmd.Context().Value(Client).(client.Client)

    projects, err := cli.ListProjects()

    if err != nil {
      log.Fatal("Unable to list projects") 
    }

    log.Print(projects)
    
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
