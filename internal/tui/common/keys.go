package common

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	Start key.Binding
	Quit  key.Binding
}

var KeyMap = keyMap{
	Start: key.NewBinding(
		key.WithKeys(" "),
		key.WithHelp("space", "start/pause/resume"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}
