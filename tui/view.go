package tui

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)


type MainModel struct { 
}

func initModel() MainModel {
  return MainModel{}
}

func (m MainModel) Init() tea.Cmd {
  return nil;
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

  switch msg := msg.(type) {
    case tea.KeyMsg:
      switch msg.String() {
        case "q":
          return m, tea.Quit
      }
  }

  return m, nil
} 


func (m MainModel) View() string {
  return "stub"
}


func Run() {
  p := tea.NewProgram(initModel())

  if _, err := p.Run(); err != nil {
    log.Fatal("erorr", err)
    os.Exit(1)
  }
}
