package cmd

import (
	"ffwizard/ffmpeg"

	"github.com/spf13/cobra"
)

var volume string

var changeVolumeCmd = &cobra.Command{
	Use:   "volume",
	Short: "change volume of your video. volume --volume",

	Run: func(cmd *cobra.Command, args []string) {
		volume, _ := cmd.Flags().GetString("volume")
		changeVolumeAction := ffmpeg.Action{Name: ffmpeg.ChangeVolume, Params: map[string]string{"Volume": volume}}
		AddAction(changeVolumeAction)
	},
}

func init() {
	changeVolumeCmd.Flags().StringVarP(&volume, "volume", "", "", "Change volume of your video")
}
