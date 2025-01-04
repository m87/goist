package cmd_test

import (
	"testing"
	"github.com/m87/goist/cmd"
)


func TestParseProjectName(t *testing.T) {
  content, project, _ := cmd.Parse("test #project")

  if content != "test" {
    t.Fatal("content", content)
  }

  if project != "project" {
    t.Fatal("project", project)
  }


  content, project, _ = cmd.Parse("test #project test")

  if content != "test test" {
    t.Fatal("content", content)
  }

  if project != "project" {
    t.Fatal("project", project)
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

  content, project, labels = cmd.Parse("test @label #project test @label2")

  if content != "test test" {
    t.Fatal("content", content)
  }

  if project != "project" {
    t.Fatal("project", project)
  }

  if len(labels) != 2 || labels[0] != "label" || labels[1] != "label2" {
    t.Fatal("labels", labels)
  }
}

