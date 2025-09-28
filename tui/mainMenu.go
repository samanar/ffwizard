package tui

import (
	"ffwizard/ffmpeg"
)

const (
	_ = iota
	ResizeStep
	RotateStep
	ConvertStep
	MuteStep
	HardSubStep
	SoftSubStep
	CompressStep
)

var MainMenuItems = []Item{
	Item{
		title:    "Resize",
		desc:     "scale video to given width and height",
		action:   ffmpeg.Action{},
		goToStep: ResizeStep,
	},
	Item{
		title:    "Rotate",
		desc:     "Rotate video",
		action:   ffmpeg.Action{},
		goToStep: RotateStep,
	},
	Item{
		title:    "Convert",
		desc:     "Change format of your video to mp4,mkv ...",
		action:   ffmpeg.Action{},
		goToStep: ConvertStep,
	},
	Item{
		title:    "Mute",
		desc:     "Mute vide",
		goToStep: MuteStep,
	},
	Item{
		title:    "Hard sub",
		desc:     "Add hard subtitle to your video",
		goToStep: HardSubStep,
	},
	Item{
		title:    "Soft sub",
		desc:     "Add soft subtitle to your video",
		goToStep: SoftSubStep,
	},
	Item{
		title:    "Compress",
		desc:     "Compress video without losing too much quality (lossy)",
		goToStep: SoftSubStep,
	},
}

var ResizeMenuItems = []Item{
	Item{
		title: "3840x2160 (4K)",
		desc:  "modern standard for TVs, cameras, YouTube, streaming platforms",
		action: ffmpeg.Action{
			Name:   ffmpeg.Resize,
			Params: map[string]string{"Width": "3840", "Height": "2160"},
		},
	},
	Item{
		title: "2560x1440 (2K)",
		desc:  "common on gaming monitors, YouTube supports it",
		action: ffmpeg.Action{
			Name:   ffmpeg.Resize,
			Params: map[string]string{"Width": "2560", "Height": "1440"},
		},
	},
	Item{
		title: "1920x1080 (1080p)",
		desc:  "Very widely used for streaming, YouTube, Blu-ray, games, and general video.",
		action: ffmpeg.Action{
			Name:   ffmpeg.Resize,
			Params: map[string]string{"Width": "1920", "Height": "1080"},
		},
	},
	Item{
		title: "1280x720 (720p)",
		desc:  "Used for smaller streams, lower bandwidth.",
		action: ffmpeg.Action{
			Name:   ffmpeg.Resize,
			Params: map[string]string{"Width": "1280", "Height": "720"},
		},
	},
	Item{
		title: "854x480 (480p)",
		desc:  "classic standard edition for mobile or low quality streams",
		action: ffmpeg.Action{
			Name:   ffmpeg.Resize,
			Params: map[string]string{"Width": "854", "Height": "480"},
		},
	},
	Item{
		title: "640x360 (360p)",
		desc:  "Very low-res, often fallback on bad connection",
		action: ffmpeg.Action{
			Name:   ffmpeg.Resize,
			Params: map[string]string{"Width": "640", "Height": "360"},
		},
	},
	Item{
		title: "426x240 (240p)",
		desc:  "mainly for old phones or extreme compression",
		action: ffmpeg.Action{
			Name:   ffmpeg.Resize,
			Params: map[string]string{"Width": "426", "Height": "240"},
		},
	},
}

var RotateMenuITems = []Item{
	Item{
		title: "90° clockwise + vertical flip",
		action: ffmpeg.Action{
			Name:   ffmpeg.Rotate,
			Params: map[string]string{"Angle": "90"},
		},
	},
	Item{
		title: "180° clockwise using two 90°",
		action: ffmpeg.Action{
			Name:   ffmpeg.Rotate,
			Params: map[string]string{"Angle": "180"},
		},
	},
	Item{
		title: "90° counter clockwise",
		action: ffmpeg.Action{
			Name:   ffmpeg.Rotate,
			Params: map[string]string{"Angle": "270"},
		},
	},
}

var ConvertMenuITems = []Item{
	Item{
		title: "mp4",
		desc:  "Generic purpose. Youtube, phones, web, pretty much everywhere",
		action: ffmpeg.Action{
			Name:   ffmpeg.Convert,
			Params: map[string]string{"Format": "mp4"},
		},
	},
	Item{
		title: "mkv",
		desc:  "supports multiple audio tracks, subtitles, lossless audio. Popular with movie rips, Plex, archival",
		action: ffmpeg.Action{
			Name:   ffmpeg.Convert,
			Params: map[string]string{"Format": "mkv"},
		},
	},
	Item{
		title: "webm",
		desc:  "Web video, HTML5 native format, optimized for streaming",
		action: ffmpeg.Action{
			Name:   ffmpeg.Convert,
			Params: map[string]string{"Format": "webm"},
		},
	},
	Item{
		title: "mov",
		desc:  "Apple QuickTime format, common for editing, sometimes larger files",
		action: ffmpeg.Action{
			Name:   ffmpeg.Convert,
			Params: map[string]string{"Format": "mov"},
		},
	},
	Item{
		title: "mpg/mpeg",
		desc:  "DVDs, broadcast TV. Large files compared to H.264",
		action: ffmpeg.Action{
			Name:   ffmpeg.Convert,
			Params: map[string]string{"Format": "mpg"},
		},
	},
	Item{
		title: "gif",
		desc:  "Memes, short loops, no audio. Heavy compared to modern web formats",
		action: ffmpeg.Action{
			Name:   ffmpeg.Convert,
			Params: map[string]string{"Format": "gif"},
		},
	},
	Item{
		title: "ogg",
		desc:  "Open-source enthusiasts, Firefox days before WebM",
		action: ffmpeg.Action{
			Name:   ffmpeg.Convert,
			Params: map[string]string{"Format": "ogg"},
		},
	},
	Item{
		title: "3gp",
		desc:  "Old mobile phones, MMS, low storage devices",
		action: ffmpeg.Action{
			Name:   ffmpeg.Convert,
			Params: map[string]string{"Format": "3gp"},
		},
	},
	Item{
		title: "avi",
		desc:  "Legacy Windows format, not ideal for modern streaming",
		action: ffmpeg.Action{
			Name:   ffmpeg.Convert,
			Params: map[string]string{"Format": "avi"},
		},
	},
	Item{
		title: "wmv",
		desc:  "Windows ecosystem, legacy. Rare now outside corporate archives",
		action: ffmpeg.Action{
			Name:   ffmpeg.Convert,
			Params: map[string]string{"Format": "wmv"},
		},
	},
	Item{
		title: "flv",
		desc:  "Old Flash Video format, used for streaming (RIP Flash).",
		action: ffmpeg.Action{
			Name:   ffmpeg.Convert,
			Params: map[string]string{"Format": "flv"},
		},
	},
}
