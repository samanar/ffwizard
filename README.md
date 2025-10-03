# FFWizard 🎬✨

A friendly FFmpeg CLI + TUI wizard built with Go

[![Go Report Card](https://goreportcard.com/badge/github.com/samanar/ffwizard)](https://goreportcard.com/report/github.com/samanar/ffwizard)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## 🌟 Features

- **🖥️ Dual Interface**: Use as a traditional CLI tool or launch an interactive TUI
- **⛓️ Command Chaining**: Combine multiple operations in a single FFmpeg run
- **🎯 Simple Syntax**: Intuitive commands that hide FFmpeg complexity
- **📊 Interactive TUI**: Step-by-step wizard with beautiful Bubble Tea interface
- **🚀 Automation Ready**: Perfect for scripts and batch processing
- **💡 Beginner Friendly**: No need to memorize complex FFmpeg flags

## 📦 Installation

### Using Go Install

```bash
go install github.com/samanar/ffwizard@latest
```

### From Source

```bash
git clone https://github.com/samanar/ffwizard.git
cd ffwizard
go build -o ffwizard
```

### Prerequisites

- **Go** 1.21+ (for building from source)
- **FFmpeg** must be installed and available in your PATH

Check if FFmpeg is installed:

```bash
ffmpeg -version
```

## 🚀 Quick Start

### Interactive TUI Mode

Simply run `ffwizard` without any arguments to launch the interactive wizard:

```bash
ffwizard
```

The TUI will guide you through:

1. Selecting your input file
2. Choosing operations (compress, resize, rotate, etc.)
3. Setting parameters for each operation
4. Specifying output file
5. Executing the FFmpeg command

### CLI Mode

Use direct commands for automation and scripting:

```bash
# Compress a video
ffwizard -i input.mp4 -o output.mp4 compress --crf 21

# Resize a video
ffwizard -i input.mp4 -o output.mp4 resize --width 1920

# Rotate a video
ffwizard -i input.mp4 -o output.mp4 rotate --angle 90

# Chain multiple operations
ffwizard -i input.mp4 -o output.mp4 \
  resize --width 720 --height 480 \
  rotate --angle 90 \
  compress --crf 23
```

## 📖 Usage

### Global Flags

```bash
ffwizard [flags] [command]
```

**Flags:**

- `-i, --input <file>` - Input video file
- `-o, --output <file>` - Output video file
- `-h, --help` - Help for ffwizard

### Available Commands

#### 🗜️ Compress

Reduce video file size while maintaining quality.

```bash
ffwizard -i input.mp4 -o output.mp4 compress [flags]
```

**Flags:**

- `--crf <value>` - Constant Rate Factor (0-51, default: 23)
  - Lower values = better quality, larger file
  - 18-28 is a good range for most videos
  - 23 is the default and provides good balance

**Examples:**

```bash
# High quality compression
ffwizard -i video.mp4 -o compressed.mp4 compress --crf 18

# Balanced compression (default)
ffwizard -i video.mp4 -o compressed.mp4 compress --crf 23

# Maximum compression
ffwizard -i video.mp4 -o compressed.mp4 compress --crf 28
```

#### 📐 Resize

Change video dimensions.

```bash
ffwizard -i input.mp4 -o output.mp4 resize [flags]
```

**Flags:**

- `--width <pixels>` - Target width
- `--height <pixels>` - Target height
- Use `-1` for auto-calculation to maintain aspect ratio

**Examples:**

```bash
# Resize to Full HD width (auto height)
ffwizard -i video.mp4 -o resized.mp4 resize --width 1920 --height -1

# Resize to specific dimensions
ffwizard -i video.mp4 -o resized.mp4 resize --width 1280 --height 720

# Resize to 4K width (auto height)
ffwizard -i video.mp4 -o resized.mp4 resize --width 3840 --height -1
```

#### 🔄 Rotate

Rotate video by specified angle.

```bash
ffwizard -i input.mp4 -o output.mp4 rotate [flags]
```

**Flags:**

- `--angle <degrees>` - Rotation angle (90, 180, 270, or custom)

**Examples:**

```bash
# Rotate 90 degrees clockwise
ffwizard -i video.mp4 -o rotated.mp4 rotate --angle 90

# Rotate 180 degrees
ffwizard -i video.mp4 -o rotated.mp4 rotate --angle 180

# Rotate 270 degrees (90 counter-clockwise)
ffwizard -i video.mp4 -o rotated.mp4 rotate --angle 270
```

### ⛓️ Chaining Commands

Combine multiple operations into a single FFmpeg execution:

```bash
# Resize, rotate, and compress
ffwizard -i input.mp4 -o output.mp4 \
  resize --width 1920 --height 1080 \
  rotate --angle 90 \
  compress --crf 21

# Create a thumbnail-sized, highly compressed video
ffwizard -i large-video.mp4 -o thumbnail.mp4 \
  resize --width 640 --height -1 \
  compress --crf 28

# Process for web delivery
ffwizard -i source.mp4 -o web-video.mp4 \
  resize --width 1280 --height 720 \
  compress --crf 25
```

## 🎨 TUI Interface

When launched without subcommands, FFWizard presents an interactive terminal UI powered by Bubble Tea:

```bash
ffwizard
```

### TUI Features:

- **📂 File Selection**: Browse and select input files
- **🎯 Operation Picker**: Choose from available operations via interactive lists
- **⚙️ Parameter Configuration**: Set values using intuitive prompts
- **👁️ Live Preview**: See the FFmpeg command before execution
- **✅ Confirmation**: Review and confirm before processing

### Navigation:

- `↑/↓` or `j/k` - Navigate lists
- `Enter` - Select option
- `Esc` or `q` - Go back/quit
- `Ctrl+C` - Exit

## 🛠️ Building From Source

### Development

```bash
# Clone the repository
git clone https://github.com/samanar/ffwizard.git
cd ffwizard

# Install dependencies
go mod download

# Build
go build -o ffwizard

# Run
./ffwizard
```

### Building for Multiple Platforms

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o ffwizard-linux-amd64

# macOS
GOOS=darwin GOARCH=amd64 go build -o ffwizard-darwin-amd64

# Windows
GOOS=windows GOARCH=amd64 go build -o ffwizard-windows-amd64.exe
```

## 📚 Examples

### Batch Processing

```bash
#!/bin/bash
# Process all MP4 files in a directory

for file in *.mp4; do
    ffwizard -i "$file" -o "processed_${file}" \
        resize --width 1920 --height -1 \
        compress --crf 23
done
```

### Create Web-Optimized Videos

```bash
ffwizard -i source.mov -o web-optimized.mp4 \
  resize --width 1920 --height 1080 \
  compress --crf 25
```

### Create Multiple Resolutions

```bash
# Create 1080p version
ffwizard -i source.mp4 -o output-1080p.mp4 \
  resize --width 1920 --height 1080 \
  compress --crf 23

# Create 720p version
ffwizard -i source.mp4 -o output-720p.mp4 \
  resize --width 1280 --height 720 \
  compress --crf 23

# Create 480p version
ffwizard -i source.mp4 -o output-480p.mp4 \
  resize --width 854 --height 480 \
  compress --crf 23
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.


## 📞 Support

- **Issues**: [GitHub Issues](https://github.com/samanar/ffwizard/issues)
- **Discussions**: [GitHub Discussions](https://github.com/samanar/ffwizard/discussions)

## 🗺️ Roadmap

- [ ] Add AI prompt for chatgpt
- [ ] Support for video filters (blur, sharpen, etc.)

---

Made with ❤️ by [Saman](https://github.com/samanar)

**⭐ If you find this project useful, please consider giving it a star!**
