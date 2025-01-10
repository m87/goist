package tui

import (
	"fmt"

	"github.com/m87/goist/model"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/m87/goist/client"
)


type ProjectsViewModel struct {
  client client.Client
  projects []model.Project
  selected int
}


func (m ProjectsViewModel) Init() tea.Cmd {
  return nil
}

func (m ProjectsViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {

  case tea.KeyMsg:
    switch msg.String() {
    case "ctrl+c", "q":
      return m, tea.Quit
    case "up", "k": 
      if m.selected > 0 {
        m.selected--
      }
    case "down", "j":
      if m.selected < len(m.projects) - 1 {
        m.selected++
      }
    }
  }

  return m, nil
}

func (m ProjectsViewModel) View() string {
  s := "Projects\n\n"

  for i, project := range m.projects {
    cursor := " "
    if m.selected == i {
      cursor = ">"
    }

    s += fmt.Sprintf("%s %s\n", cursor, project.Name)
  }

  return s
}
