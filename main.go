package main

import (
	"ffwizard/cmd"
)

func main() {
	// resizeAction := ffmpeg.Action{Name: ffmpeg.Resize, Params: map[string]string{"Width": "1920", "Height": "1080"}}
	// formatAction := ffmpeg.Action{Name: ffmpeg.Convert, Params: map[string]string{"Format": "mp4"}}
	// addSoftSubAction := ffmpeg.Action{Name: ffmpeg.AddHardSub, Params: map[string]string{"Subtitle": "hard.srt"}}
	// addHardSubAction := ffmpeg.Action{Name: ffmpeg.AddSoftSub, Params: map[string]string{"Subtitle": "soft.srt"}}
	// compressAction := ffmpeg.Action{Name: ffmpeg.Compress, Params: map[string]string{"Subtitle": "soft.srt"}}
	// ffmpeg.RunFFmpeg("input.mp4", "output.mp4", []ffmpeg.Action{resizeAction, formatAction, addSoftSubAction, addHardSubAction, compressAction})
	cmd.Execute()
}
