package cmd

import (
	"ffwizard/ffmpeg"

	"github.com/spf13/cobra"
)

var angle string

func NormalizeAngle(angle string) string {
	switch angle {
	case "90", "180", "270", "360":
		return angle
	default:
		return "90"
	}
}

var RotateCmd = &cobra.Command{
	Use:   "rotate",
	Short: "Rotate your video. Usage: --angle [90, 180 , 270, 360]",
	Run: func(cmd *cobra.Command, args []string) {
		angle, _ := cmd.Flags().GetString("angle")
		rotateAction := ffmpeg.Action{Name: ffmpeg.Rotate, Params: map[string]string{"Angle": NormalizeAngle(angle)}}
		AddAction(rotateAction)
	},
}

func init() {
	RotateCmd.Flags().StringVarP(&angle, "angle", "", "", "Rotation angle")
}
