package tui

import (
	"errors"
	"ffwizard/ffmpeg"
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func getModelListFromStep(step uint) (list.Model, error) {
	switch step {
	case 0:
		return GetList("FFWIZARD", MainMenuItems), nil
	case ResizeStep:
		return GetList("Resize", ResizeMenuItems), nil
	case RotateStep:
		return GetList("Rotate", RotateMenuITems), nil
	case ConvertStep:
		return GetList("Rotate", ConvertMenuITems), nil
	default:
		return list.Model{}, errors.New("invalid step")
	}

}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(Item)
			if ok {
				if i.action.Name != 0 {
					m.AddAction(i.action)
				}
				oldStep := m.step
				m.step = i.goToStep
				fmt.Println(m.step)
				switch m.step {
				case 0, ConvertStep, CompressStep, RotateStep, ResizeStep:
					newList, err := getModelListFromStep(m.step)
					if err != nil {
						m.list = newList
						return m, nil
					} else {
						return m, tea.Quit
					}
				case HardSubStep:
					if oldStep == 0 {
						var fileName = m.textInput.Value()
						fmt.Println(fileName)
						if fileName != "" {
							m.AddAction(ffmpeg.Action{Name: ffmpeg.AddHardSub, Params: map[string]string{"Subtitle": fileName}})
							m.step = 0
							m.list, _ = getModelListFromStep(0)
							return m, nil
						}
					}
					m.textInput = GetInput("Enter name of your hard subtitle file: Ex: hardsub.srt")
					return m, nil

				default:
					return m, tea.Quit
				}
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	switch m.step {
	case HardSubStep, SoftSubStep:
		m.textInput, cmd = m.textInput.Update(msg)
	default:
		m.list, cmd = m.list.Update(msg)
	}
	return m, cmd
}
