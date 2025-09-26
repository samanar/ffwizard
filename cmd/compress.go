/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"ffwizard/ffmpeg"

	"github.com/spf13/cobra"
)

var compressCmd = &cobra.Command{
	Use:   "compress",
	Short: "compress your video",
	Run: func(cmd *cobra.Command, args []string) {
		compressAction := ffmpeg.Action{Name: ffmpeg.Compress}
		AddAction(compressAction)
	},
}

func init() {
	rootCmd.AddCommand(compressCmd)
}
