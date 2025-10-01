package tui

import (
	"ffwizard/ffmpeg"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

var tuiActions []ffmpeg.Action

type Model struct {
	list      list.Model
	textInput textinput.Model
	step      uint
	quitting  bool
}

func (m *Model) AppendAction(a ffmpeg.Action) {
	tuiActions = append(tuiActions, a)
}

func (m *Model) AddAction(newAction ffmpeg.Action) {
	for i, a := range tuiActions {
		if a.Name == newAction.Name {
			tuiActions[i] = newAction
			return
		}
	}
	m.AppendAction(newAction)
}

func InitialModel() Model {
	return Model{
		step: 0,
		list: GetList("ffwizard", MainMenuItems),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
