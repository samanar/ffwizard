package ffmpeg

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type ActionName int

const (
	_ ActionName = iota
	Resize
	Convert
	Rotate
	Mute
	NormalizeAudio
	ChangeVolume
	ExtractAudio
	ReplaceAudio
	AddHardSub
	AddSoftSub
	CRF
	VideoBitrate
)

type Action struct {
	Name   ActionName
	Params map[string]string
}

func getFormatParams(format string) (videoCodec, audioCodec, extra string, ok bool) {
	switch format {
	case "mp4":
		return "libx264", "aac", "-crf 23 -preset fast -b:a 192k", true
	case "mkv":
		return "libx264", "aac", "-crf 23 -preset fast -b:a 192k", true
	case "webm":
		return "libvpx-vp9", "libopus", "-crf 30 -b:v 0", true
	case "avi":
		return "mpeg4", "libmp3lame", "-q:v 5 -b:a 192k", true
	case "mov":
		return "libx264", "aac", "-crf 23 -preset fast -b:a 192k", true
	case "flv":
		return "libx264", "libmp3lame", "-crf 28 -preset fast -b:a 128k", true
	case "ogg":
		return "libtheora", "libvorbis", "-q:v 7 -q:a 5", true
	case "wmv":
		return "wmv2", "wmav2", "-b:v 2000k -b:a 192k", true
	case "3gp":
		return "libx264", "aac", "-crf 28 -preset ultrafast -b:a 96k -s 352x288", true
	case "mpg", "mpeg":
		return "mpeg2video", "mp2", "-q:v 5 -b:a 192k", true
	case "gif":
		return "gif", "", "", true
	case "ts":
		return "mpeg2video", "mp2", "-q:v 5 -b:a 192k", true
	case "m4v":
		return "libx264", "aac", "-crf 23 -preset fast -b:a 160k", true
	default:
		return "", "", "", false
	}
}

func getFormat(filename string, actions []Action) string {
	for i := len(actions) - 1; i >= 0; i-- {
		if actions[i].Name == Convert {
			return actions[i].Params["Format"]
		}
	}
	ext := filepath.Ext(filename) // e.g. ".mp4"
	if ext == "" {
		return ""
	}
	return strings.ToLower(strings.TrimPrefix(ext, "."))
}

func getSubTitleEncoding(format string) (string, bool) {
	switch format {
	case "mp4":
		return "mov_text", true

	case "mkv", "webm", "ogg", "ogv":
		return "copy", true

	case "avi", "flv", "ts", "mpegts":
		return "", false

	default:
		return "copy", true
	}
}

func getTransposeFromAngle(angle int) (string, error) {
	switch angle {
	case 90:
		return "transpose=1", nil // 90° clockwise
	case 180:
		// Rotate 180° using two 90° rotates
		return "transpose=1,transpose=1", nil
	case 270:
		return "transpose=2", nil // 90° counterclockwise
	case 0:
		return "", nil // no rotation needed
	default:
		return "", errors.New("unsupported angle: must be 0, 90, 180, or 270")
	}
}

func buildCommand(input, output string, actions []Action) []string {
	args := []string{"-i", input}
	for _, action := range actions {
		switch action.Name {

		case Resize:
			args = append(args, "-vf", fmt.Sprintf("scale=%s:%s", action.Params["Width"], action.Params["Height"]))

		case Convert:
			format := action.Params["Format"]
			v, a, extra, ok := getFormatParams(format)
			if !ok {
				fmt.Printf("Format %s not supported\n", format)
				os.Exit(1)
			}
			args = append(args, fmt.Sprintf("-c:v %s -c:a %s %s", v, a, extra))

		case AddHardSub:
			args = append(args, fmt.Sprintf("-vf subtitle=%s", action.Params["Subtitle"]))

		case AddSoftSub:
			subTitleEncoding, ok := getSubTitleEncoding(getFormat(output, actions))
			if ok {
				args = append(args, fmt.Sprintf("-i %s -c:s %s", action.Params["Subtitle"], subTitleEncoding))
			}

		case Mute:
			args = append(args, "-an")

		case ChangeVolume:
			volumeParameter := action.Params["Volume"]
			if volumeParameter == "" {
				volumeParameter = "2.0"
			}
			args = append(args, fmt.Sprintf("-af volume %s", volumeParameter))

		case ExtractAudio:
			args = append(args, "-q:a 0 -map a")

		case NormalizeAudio:
			args = append(args, "-af loudnorm")

		case ReplaceAudio:
			replace := action.Params["Audio"]
			if replace == "" {
				replace = "newAudio.mp3"
			}
			args = append(args, fmt.Sprintf("-i %s -c:v copy -map 0:v:0 -map 1:a:0", replace))

		case Rotate:
			angleStr := action.Params["Angle"]
			angle, err := strconv.Atoi(angleStr)
			if err != nil {
				fmt.Println("Invalid angle:", err)
				os.Exit(1)
			}
			transpose, err := getTransposeFromAngle(angle)
			if err != nil {
				fmt.Println("Error getting transpose:", err)
				os.Exit(1)
			}
			args = append(args, fmt.Sprintf("-vf transpose=%s", transpose))

		case CRF:
			crf := action.Params["Crf"]
			if crf == "" {
				crf = "23"
			}
			args = append(args, fmt.Sprintf("-vcodec libx264 -crf %s", crf))

		case VideoBitrate:
			VideoBitrate := action.Params["VideoBitrate"]
			if VideoBitrate == "" {
				VideoBitrate = "2000k"
			}
			args = append(args, fmt.Sprintf("-b:v %s -b:a 128k", VideoBitrate))

		}

	}
	args = append(args, output)
	return args
}

func RunFFmpeg(input, output string, actions []Action) error {
	if input == "" {
		input = "sample-input.mp4"
	}
	if output == "" {
		output = "output.mp4"
	}
	args := buildCommand(input, output, actions)
	fmt.Println("Your command is:")
	fmt.Println("ffmpeg " + strings.Join(args, " "))

	// Ask for confirmation
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Do you want to run this command? (y/n): ")
	confirm, _ := reader.ReadString('\n')
	confirm = strings.TrimSpace(confirm)

	if confirm == "y" || confirm == "Y" {
		// Run the command
		cmd := exec.Command("ffmpeg", args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			fmt.Println("Error running ffmpeg:", err)
			return err
		} else {
			fmt.Println("FFmpeg finished successfully!")
		}
	}
	return nil

}
