package tui

func (m Model) View() string {
	if m.quitting {
		return "Exiting wizard...\n"
	}

	switch m.step {
	case HardSubStep, SoftSubStep:
		return m.textInput.View()
	default:
		return m.list.View()
	}

	return `
Action added. 
Press another number to add more, or press enter to finish.
`
}
