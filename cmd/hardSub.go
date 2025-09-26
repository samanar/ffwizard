package cmd

import (
	"ffwizard/ffmpeg"

	"github.com/spf13/cobra"
)

var subtitle string

// hardSubCmd represents the hardSub command
var hardSubCmd = &cobra.Command{
	Use:   "hardSub",
	Short: "Add hard subtitle to your video. Usage: --subtitle hard-subtitle.srt",

	Run: func(cmd *cobra.Command, args []string) {
		subtitle, _ := cmd.Flags().GetString("subtitle")
		if subtitle == "" {
			subtitle = "hard-subtitle.srt"
		}
		hardSubAction := ffmpeg.Action{Name: ffmpeg.AddHardSub, Params: map[string]string{"Subtitle": subtitle}}
		AddAction(hardSubAction)
	},
}

func init() {
	hardSubCmd.Flags().StringVarP(&subtitle, "subtitle", "", "", "adding subtitle")
}
