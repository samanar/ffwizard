package tui

import (
	"errors"
	"ffwizard/ffmpeg"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func getModelListFromStep(m Model) (list.Model, error) {
	switch m.step {
	case 0:
		menuItems := MainMenuItems
		if len(tuiActions) > 0 {
			menuItems = append(menuItems, Item{title: "Create command", desc: "create your final command", goToStep: FinalStep})
			menuItems = append([]list.Item{Item{title: "Create command", desc: "create your final command", goToStep: FinalStep}}, menuItems...)
		}
		return GetList("ffwizard", menuItems), nil
	case ResizeStep:
		return GetList("Resize", ResizeMenuItems), nil
	case RotateStep:
		return GetList("Rotate", RotateMenuITems), nil
	case ConvertStep:
		return GetList("Convert to", ConvertMenuITems), nil
	case SoundStep:
		return GetList("Sound optionw", SoundMeuItems), nil
	case ChangeVolumeStep:
		return GetList("Change volume to", ChangeVolumeMenuItems), nil
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
				m.step = i.goToStep
				switch m.step {
				case 0, ConvertStep, CompressStep, RotateStep, ResizeStep, SoundStep, ChangeVolumeStep:
					newList, err := getModelListFromStep(m)
					if err == nil {
						m.list = newList
						return m, nil
					} else {
						return m, tea.Quit
					}

				case HardSubStep, SoftSubStep:
					var fileName = m.textInput.Value()
					if fileName != "" {
						actionName := ffmpeg.AddHardSub
						if m.step == SoftSubStep {
							actionName = ffmpeg.AddSoftSub
						}
						m.AddAction(ffmpeg.Action{Name: actionName, Params: map[string]string{"Subtitle": fileName}})
						m.step = 0
						m.textInput.Reset()
						m.list, _ = getModelListFromStep(m)
						return m, nil
					}
					m.textInput = GetInput("Enter name of your  subtitle file: Ex: subtitle.srt")
					return m, nil

				case ReplaceAudioStep:
					var fileName = m.textInput.Value()
					if fileName != "" {
						m.AddAction(ffmpeg.Action{Name: ffmpeg.ReplaceAudio, Params: map[string]string{"Audio": fileName}})
						m.step = 0
						m.textInput.Reset()
						m.list, _ = getModelListFromStep(m)
						return m, nil
					}
					m.textInput = GetInput("Enter name of your replace audio file: Ex: replace.mp3")
					return m, nil

				case FinalStep:
					return m, tea.Quit

				default:
					return m, tea.Quit
				}
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	switch m.step {
	case HardSubStep, SoftSubStep, ReplaceAudioStep:
		m.textInput, cmd = m.textInput.Update(msg)
	default:
		m.list, cmd = m.list.Update(msg)
	}
	return m, cmd
}
