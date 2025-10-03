package cmd

import (
	"ffwizard/ffmpeg"

	"github.com/spf13/cobra"
)

var crf, videoBitrate string
var CompressCmd = &cobra.Command{
	Use:   "compress",
	Short: "compress your video",
	Run: func(cmd *cobra.Command, args []string) {
		crf, _ := cmd.Flags().GetString("crf")
		videoBitrate, _ := cmd.Flags().GetString("bitrate")
		crfAction := ffmpeg.Action{Name: ffmpeg.CRF, Params: map[string]string{"Crf": crf}}
		bitrateAction := ffmpeg.Action{Name: ffmpeg.VideoBitrate, Params: map[string]string{"VideoBitrate": videoBitrate}}
		AddAction(crfAction)
		AddAction(bitrateAction)
	},
}

func init() {
	CompressCmd.Flags().StringVarP(&crf, "crf", "", "", "Change video crf")
	CompressCmd.Flags().StringVarP(&videoBitrate, "bitrate", "", "", "Change video bitrate")
}
