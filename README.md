# 🎬 ffmpeg-tool

A friendly **FFmpeg CLI + TUI wizard** built with **Go**, **Cobra**, and **Bubbletea**.  
Use it in two ways:

1. **Direct CLI commands** like `resize`, `rotate`, `compress` — perfect for automation.
2. **Interactive TUI wizard** if no subcommand is given — step by step FFmpeg command builder.

---

## 🚀 Features

- Chain multiple actions into a single ffmpeg run:
  ```bash
  ffmpeg-tool -i input.mp4 -o output.mp4 resize --width 720 --height 480 rotate --angle 90 compress 
  ```
