/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"ffwizard/ffmpeg"

	"github.com/spf13/cobra"
)

var width, height string

// resizeCmd represents the resize command
var resizeCmd = &cobra.Command{
	Use:   "resize",
	Short: "Resize your video. Usage : --width 1920 --height 1080",

	Run: func(cmd *cobra.Command, args []string) {
		width, _ := cmd.Flags().GetString("width")
		height, _ := cmd.Flags().GetString("height")
		if width == "" {
			width = "-1"
		}
		if height == "" {
			height = "-1"
		}
		resizeAction := ffmpeg.Action{Name: ffmpeg.Resize, Params: map[string]string{"Width": width, "Height": height}}
		AddAction(resizeAction)
	},
}

func init() {
	resizeCmd.Flags().StringVarP(&width, "width", "", "", "Target width")
	resizeCmd.Flags().StringVarP(&height, "height", "", "", "Target height")
}
