package tui

import (
	"ffwizard/ffmpeg"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	list     list.Model
	step     int
	choice   string
	actions  []ffmpeg.Action
	quitting bool
}

func InitialModel() Model {
	return Model{
		step:    0,
		list:    GetList("ffwizard", MainMenuItems),
		actions: []ffmpeg.Action{},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
