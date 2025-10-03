package cmd

import (
	"ffwizard/ffmpeg"

	"github.com/spf13/cobra"
)

// muteCmd represents the mute command
var ExtractAudioCmd = &cobra.Command{
	Use:   "extract-audio",
	Short: "extract audio of your video",

	Run: func(cmd *cobra.Command, args []string) {
		extractAudioAction := ffmpeg.Action{Name: ffmpeg.ExtractAudio}
		AddAction(extractAudioAction)
	},
}
