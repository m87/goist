/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// taskCmd represents the task command
var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Add new task",
	Long: `Add new task. E.g. 
  goist create task "new task #porject @label"
  goist create task "new task" -p project -l label`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("task called")
    fmt.Println(viper.Get("token"))
	},
}

func init() {
	createCmd.AddCommand(taskCmd)
  taskCmd.Flags().String("project", "p", "Parent project")
  taskCmd.Flags().String("label", "l", "Tasks label")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// taskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// taskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
