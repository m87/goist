package tui

import (
  
  tea "github.com/charmbracelet/bubbletea"
)


type TasksViewModel struct {
}


func (m TasksViewModel) Init() tea.Cmd {
  return nil
}

func (m TasksViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  return m, nil
}

func (m TasksViewModel) View() string {
  return "tasks stub"
}
