package models

import (
	"time"

	"github.com/charmbracelet/bubbles/stopwatch"
)

type Mode int

const (
	Idle Mode = iota
	Working
	OnBreak
)

type MainModel struct {
	Stopwatch  stopwatch.Model
	Mode       Mode
	BreakTime  time.Duration
	BreakStart time.Time
	Project    string
}

func NewMainModel(project string) MainModel {
	return MainModel{
		Stopwatch: stopwatch.New(),
		Mode:      Idle,
		Project:   project,
	}
}
