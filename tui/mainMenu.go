package tui

import (
	"ffwizard/ffmpeg"

	"github.com/charmbracelet/bubbles/list"
)

var MainMenuItems = []list.Item{
	Item{
		title:  "Resize video",
		desc:   "scale video to given width and height",
		action: ffmpeg.Action{},
	},
	Item{
		title:  "Rotate video",
		desc:   "Rotate video",
		action: ffmpeg.Action{},
	},
}
