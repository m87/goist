package tui

import (
  "github.com/m87/goist/model"
  
  tea "github.com/charmbracelet/bubbletea"
)


type ProjectsViewModel struct {
  projects []model.Project
}


func (m ProjectsViewModel) Init() tea.Cmd {
  return nil
}

func (m ProjectsViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  return m, nil
}

func (m ProjectsViewModel) View() string {
  return "projects stub"
}
