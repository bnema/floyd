package updates

import (
	"time"

	"github.com/bnema/floyd/internal/tui/common"
	"github.com/bnema/floyd/internal/tui/models"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func UpdateMain(msg tea.Msg, m models.MainModel) (models.MainModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, common.KeyMap.Quit):
			return m, tea.Quit
		case key.Matches(msg, common.KeyMap.Start):
			switch m.Mode {
			case models.Idle:
				m.Mode = models.Working
				return m, m.Stopwatch.Start()
			case models.Working:
				m.Mode = models.OnBreak
				flowTime := m.Stopwatch.Elapsed()
				m.BreakTime = flowTime / 4
				m.BreakStart = time.Now()
				m.Stopwatch.Stop()
			case models.OnBreak:
				m.Mode = models.Working
				m.Stopwatch.Reset()
				return m, m.Stopwatch.Start()
			}
		}
	}

	var cmd tea.Cmd
	m.Stopwatch, cmd = m.Stopwatch.Update(msg)
	return m, cmd
}
