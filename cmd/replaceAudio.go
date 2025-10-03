package cmd

import (
	"ffwizard/ffmpeg"

	"github.com/spf13/cobra"
)

var replaceAudio string

var replaceAudioCmd = &cobra.Command{
	Use:   "replace-audio",
	Short: "replace audio of your video. replace-audio --audio",

	Run: func(cmd *cobra.Command, args []string) {
		replaceAudio, _ := cmd.Flags().GetString("audio")
		extractAudioAction := ffmpeg.Action{Name: ffmpeg.ReplaceAudio, Params: map[string]string{"Audio": replaceAudio}}
		AddAction(extractAudioAction)
	},
}

func init() {
	replaceAudioCmd.Flags().StringVarP(&replaceAudio, "audio", "", "", "Replace audio from your video with given")
}
