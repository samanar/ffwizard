package cmd

import (
	"ffwizard/ffmpeg"

	"github.com/spf13/cobra"
)

var softSub string

// hardSubCmd represents the hardSub command
var softSubCmd = &cobra.Command{
	Use:   "softSub",
	Short: "Add soft subtitle to your video. Usage: --subtitle soft-subtitle.srt",

	Run: func(cmd *cobra.Command, args []string) {
		softSub, _ := cmd.Flags().GetString("subtitle")
		if softSub == "" {
			softSub = "soft-subtitle.srt"
		}
		action := ffmpeg.Action{Name: ffmpeg.AddSoftSub, Params: map[string]string{"Subtitle": subtitle}}
		AddAction(action)
	},
}

func init() {
	softSubCmd.Flags().StringVarP(&softSub, "subtitle", "", "", "adding subtitle")
}
