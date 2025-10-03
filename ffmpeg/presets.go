package ffmpeg

import "fmt"

type PresetType int

const (
	YoutubeFHD PresetType = iota
	Youtube4k
	InstagramFeed
	InstagramStory
	TikTok
	X
	WhatsApp
	Facebook
	LinkedIn
	SnapChat
	Web
	Gif
	LosslessArchive
)

var Presets = []PresetType{Youtube4k, YoutubeFHD, InstagramFeed, InstagramStory, TikTok, Facebook, WhatsApp, X, Gif, Web, LinkedIn, SnapChat, LosslessArchive}

func (p PresetType) String() string {
	switch p {
	case YoutubeFHD:
		return "Youtube FHD"
	case Youtube4k:
		return "Youtube 4K"
	case InstagramFeed:
		return "Instagram Feed"
	case InstagramStory:
		return "Instagram story/reel"
	case TikTok:
		return "TikTok"
	case X:
		return "X/Twitter"
	case WhatsApp:
		return "WhatsApp"
	case Facebook:
		return "Facebook"
	case LinkedIn:
		return "LinkedIn"
	case SnapChat:
		return "SnapChat"
	case Web:
		return "Web"
	case Gif:
		return "Gif"
	case LosslessArchive:
		return "Lossless archive"
	default:
		return ""
	}
}

func ParsePreset(s string) (PresetType, error) {
	for _, p := range Presets {
		if p.String() == s {
			return p, nil
		}
	}
	return -1, fmt.Errorf("unknown preset")
}

func (p PresetType) Command() (string, error) {
	switch p {
	case YoutubeFHD:
		return "-c:v libx264 -preset slow -crf 22 -maxrate 8M -bufsize 16M -c:a aac -b:a 128k -movflags +faststart", nil
	case Youtube4k:
		return "-c:v libx264 -preset slow -crf 20 -pix_fmt yuv420p -c:a aac -b:a 192k -ar 48000 -movflags +faststart", nil
	case InstagramFeed:
		return "-vf 'scale=1080:1080:force_original_aspect_ratio=decrease,pad=1080:1080:(ow-iw)/2:(oh-ih)/2' -c:v libx264 -preset fast -crf 23 -maxrate 5M  ,nil-bufsize10M -c:a aac -b:a 128k", nil
	case InstagramStory:
		return "-vf 'scale=1080:1920:force_original_aspect_ratio=decrease,pad=1080:1920:(ow-iw)/2:(oh-ih)/2' -c:v libx264 -preset fast -crf 23 -maxrate 5M  ,nil-bufsize10M -c:a aac -b:a 128k", nil
	case TikTok:
		return "-vf 'scale=1080:1920:force_original_aspect_ratio=decrease,pad=1080:1920:(ow-iw)/2:(oh-ih)/2' -c:v libx264 -preset veryfast -crf 28 -maxrate  ,nil4M-bufsize 8M -c:a aac -b:a 128k", nil
	case X:
		return "-vf 'scale=1280:720:force_original_aspect_ratio=decrease' -c:v libx264 -preset fast -crf 28 -maxrate 3M -bufsize 6M -c:a aac -b:a 128k", nil
	case WhatsApp:
		return "-vf 'scale=640:480:force_original_aspect_ratio=decrease' -c:v libx264 -preset ultrafast -crf 30 -maxrate 1M -bufsize 2M -c:a aac -b:a 96k", nil
	case Facebook:
		return "-vf 'scale=1280:720:force_original_aspect_ratio=decrease' -c:v libx264 -preset fast -crf 25 -maxrate 4M -bufsize 8M -c:a aac -b:a 128k", nil
	case LinkedIn:
		return "-vf 'scale=1920:1080:force_original_aspect_ratio=decrease' -c:v libx264 -preset fast -crf 24 -maxrate 6M -bufsize 12M -c:a aac -b:a 128k", nil
	case SnapChat:
		return "-vf 'scale=1080:1920:force_original_aspect_ratio=decrease,pad=1080:1920:(ow-iw)/2:(oh-ih)/2' -c:v libx264 -preset fast -crf 25 -maxrate 4M bufsize8M -c:a aac -b:a 128k", nil
	case Web:
		return "-c:v libx264 -preset fast -crf 26 -movflags +faststart -c:a aac -b:a 128k", nil
	case Gif:
		return "-vf 'fps=10,scale=320:-1:flags=lanczos' -t 10", nil
	case LosslessArchive:
		return "-c:v libx264 -preset veryslow -crf 0 -c:a flac", nil
	default:
		return "", fmt.Errorf("invalid preset1")
	}

}
