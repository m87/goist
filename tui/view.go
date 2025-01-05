package tui

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type Context int

const (
  PROJECTS Context = iota
  TASKS
  
)


type MainModel struct { 
  context Context

  projectsModel ProjectsViewModel
  tasksModel TasksViewModel
}

func initModel() MainModel {
  return MainModel{
    projectsModel: ProjectsViewModel{},
    tasksModel: TasksViewModel{},
  }
}

func (m MainModel) Init() tea.Cmd {
  return nil;
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch m.context {
  case PROJECTS:
    subModel, cmd := m.projectsModel.Update(msg)
    m.projectsModel = subModel.(ProjectsViewModel)
    if cmd != nil {
      return m, cmd
    }
  }

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
