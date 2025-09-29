package tui

import (
	"ffwizard/ffmpeg"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	list      list.Model
	textInput textinput.Model
	step      uint
	actions   []ffmpeg.Action
	quitting  bool
}

func (m *Model) AppendAction(a ffmpeg.Action) {
	m.actions = append(m.actions, a)
}

func (m *Model) AddAction(newAction ffmpeg.Action) {
	for i, a := range m.actions {
		if a.Name == newAction.Name {
			m.actions[i] = newAction
			return
		}
	}
	m.AppendAction(newAction)
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
