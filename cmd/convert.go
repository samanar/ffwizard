package cmd

import (
	"ffwizard/ffmpeg"

	"github.com/spf13/cobra"
)

var format string

var convertCommand = &cobra.Command{
	Use:   "convert",
	Short: "Convert your video to a new format like mp4, mkv , ...",
	Run: func(cmd *cobra.Command, args []string) {
		format, _ := cmd.Flags().GetString("format")
		formatAction := ffmpeg.Action{Name: ffmpeg.Convert, Params: map[string]string{"Format": format}}
		AddAction(formatAction)
	},
}

func init() {
	convertCommand.Flags().StringVarP(&format, "format", "", "", "Video format")
}
