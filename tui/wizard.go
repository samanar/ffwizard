package tui

import (
	"ffwizard/ffmpeg"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// RunWizard launches the Bubbletea wizard
func RunWizard() ([]ffmpeg.Action, error) {
	p := tea.NewProgram(InitialModel())
	m, err := p.Run()
	if err != nil {
		return nil, err
	}

	_, ok := m.(Model)
	if !ok {
		return nil, fmt.Errorf("unexpected model type")
	}

	if len(tuiActions) == 0 {
		fmt.Println("No actions selected. Exiting.")
		os.Exit(0)
	}

	return tuiActions, nil
}
