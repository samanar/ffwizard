package tui

func (m Model) View() string {
	if m.quitting {
		return "Exiting wizard...\n"
	}

	if m.step == 0 {
		return m.list.View()
	}

	return `
Action added. 
Press another number to add more, or press enter to finish.
`
}
