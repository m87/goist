package cmd_test

import (
	"testing"

	"github.com/m87/goist/cmd"
)


func TestParseProjectName(t *testing.T) {
  content, project, labels := cmd.Parse("test #project")

  if content != "test" {
    t.Fatal("content", content)
  }

  if project != "project" {
    t.Fatal("project", project)
  }


  content, project, labels = cmd.Parse("test #project test")

  if content != "test test" {
    t.Fatal("content", content)
  }

  if project != "project" {
    t.Fatal("project", project)
  }

  if labels == nil {

  }
}



func TestParseLabel(t *testing.T) {
  content, project, labels := cmd.Parse("test #project @label")

  if content != "test" {
    t.Fatal("content", content)
  }

  if project != "project" {
    t.Fatal("project", project)
  }

  if len(labels) == 0 {
    t.Fatal("labels")
  }

  content, project, labels = cmd.Parse("test @label #project test")

  if content != "test test" {
    t.Fatal("content", content)
  }

  if project != "project" {
    t.Fatal("project", project)
  }
}

