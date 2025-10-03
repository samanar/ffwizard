package cmd

import (
	"ffwizard/ffmpeg"

	"github.com/spf13/cobra"
)

// muteCmd represents the mute command
var NormalizeAudioCmd = &cobra.Command{
	Use:   "normalize-audio",
	Short: "normalize audio of your video",

	Run: func(cmd *cobra.Command, args []string) {
		NormalizeAudioAction := ffmpeg.Action{Name: ffmpeg.NormalizeAudio}
		AddAction(NormalizeAudioAction)
	},
}
