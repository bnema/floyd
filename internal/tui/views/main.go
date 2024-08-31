package views

import (
	"fmt"
	"time"

	"github.com/bnema/floyd/internal/tui/common"
	"github.com/bnema/floyd/internal/tui/models"
	"github.com/bnema/floyd/internal/version"
)

func ViewMain(m models.MainModel) string {
	var status, timeDisplay, instructions string

	switch m.Mode {
	case models.Idle:
		status = "Ready"
		timeDisplay = "00:00:00"
		instructions = "Press SPACE to start working"
	case models.Working:
		status = "Working"
		timeDisplay = m.Stopwatch.View()
		instructions = "Press SPACE to take a break"
	case models.OnBreak:
		status = "Break"
		elapsed := time.Since(m.BreakStart)
		remaining := m.BreakTime - elapsed
		if remaining < 0 {
			remaining = 0
		}
		timeDisplay = remaining.Round(time.Second).String()
		instructions = "Press SPACE to resume work"
	}

	return common.AppStyle.Render(fmt.Sprintf(
		"%s\n\n%s: %s\n\n%s\nProject: %s\nPress q to quit",
		common.TitleStyle.Render("Floyd - "+version.PrintVersion()),
		status,
		timeDisplay,
		instructions,
		m.Project,
	))
}
