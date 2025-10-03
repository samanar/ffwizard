package cmd

import (
	"ffwizard/ffmpeg"

	"github.com/spf13/cobra"
)

// muteCmd represents the mute command
var muteCmd = &cobra.Command{
	Use:   "mute",
	Short: "mute you video",

	Run: func(cmd *cobra.Command, args []string) {
		muteAction := ffmpeg.Action{Name: ffmpeg.Mute}
		AddAction(muteAction)
	},
}
