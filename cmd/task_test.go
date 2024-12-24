package cmd_test

import (
	"testing"

	"github.com/m87/goist/cmd"
)


func TestParseTask(t *testing.T) {
  content, project := cmd.Parse("test #project")

  if content != "test" {
    t.Fatal("content", content)
  }

  if project != "project" {
    t.Fatal("project", project)
  }

}
