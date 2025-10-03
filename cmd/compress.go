package cmd

import (
	"ffwizard/ffmpeg"

	"github.com/spf13/cobra"
)

var CompressCmd = &cobra.Command{
	Use:   "compress",
	Short: "compress your video",
	Run: func(cmd *cobra.Command, args []string) {
		compressAction := ffmpeg.Action{Name: ffmpeg.Compress}
		AddAction(compressAction)
	},
}

func init() {
	rootCmd.AddCommand(CompressCmd)
}
