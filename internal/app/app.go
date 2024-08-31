package app

import (
	"github.com/bnema/floyd/internal/tui/models"
	"github.com/bnema/floyd/internal/tui/updates"
	"github.com/bnema/floyd/internal/tui/views"
	tea "github.com/charmbracelet/bubbletea"
)

type floydModel struct {
	models.MainModel
}

func (m floydModel) Init() tea.Cmd {
	return nil
}

func (m floydModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	newModel, cmd := updates.UpdateMain(msg, m.MainModel)
	m.MainModel = newModel
	return m, cmd
}

func (m floydModel) View() string {
	return views.ViewMain(m.MainModel)
}

func Run(project string) error {
	initialModel := floydModel{models.NewMainModel(project)}
	p := tea.NewProgram(initialModel, tea.WithAltScreen())
	_, err := p.Run()
	return err
}
