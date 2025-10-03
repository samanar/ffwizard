package tui

import (
	"ffwizard/ffmpeg"
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// RunWizard launches the Bubbletea wizard
func RunWizard() ([]ffmpeg.Action, error) {
	p := tea.NewProgram(InitialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}

	if len(tuiActions) == 0 {
		fmt.Println("No actions selected. Exiting.")
		os.Exit(0)
	}

	return tuiActions, nil
}
