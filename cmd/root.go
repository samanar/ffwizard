package cmd

import (
	"ffwizard/ffmpeg"
	"ffwizard/tui"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var input, output string

var actions []ffmpeg.Action

func AppendAction(a ffmpeg.Action) {
	actions = append(actions, a)
}

func AddAction(newAction ffmpeg.Action) {
	for i, a := range actions {
		if a.Name == newAction.Name {
			actions[i] = newAction
			return
		}
	}
	AppendAction(newAction)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ffwizard",
	Short: "Interactive CLI to build and run FFmpeg commands",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		actions, err = tui.RunWizard()
		if err != nil {
			fmt.Println("Error in TUI:", err)
			os.Exit(1)
		}
		fmt.Println("No subcommands provided")
	},
}

func init() {
	// Global persistent flags
	rootCmd.PersistentFlags().StringVarP(&input, "input", "i", "", "Input video file")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "Output video file")

	rootCmd.AddCommand(resizeCmd, rotateCmd, compressCmd, hardSubCmd, softSubCmd, convertCommand, muteCmd)
}

func Execute() {
	// Parse global flags first
	rootCmd.ParseFlags(os.Args[1:])
	// args := rootCmd.Flags().Args() // only positional args (subcommands)
	args := os.Args[1:]

	subcommands := rootCmd.Commands()

	for i := 0; i < len(args); {
		found := false
		for _, sub := range subcommands {
			if args[i] == sub.Use {
				// collect args for this subcommand
				j := i + 1
				subArgs := []string{}
				for j < len(args) {
					isNext := false
					for _, s := range subcommands {
						if args[j] == s.Use {
							isNext = true
							break
						}
					}
					if isNext {
						break
					}
					subArgs = append(subArgs, args[j])
					j++
				}

				sub.SetArgs(subArgs)
				if err := sub.ParseFlags(subArgs); err != nil {
					fmt.Println("Error parsing flags for", sub.Use, ":", err)
				}
				if sub.RunE != nil {
					_ = sub.RunE(sub, sub.Flags().Args())
				} else if sub.Run != nil {
					sub.Run(sub, sub.Flags().Args())
				}

				i = j
				found = true
				break
			}
		}

		if !found {
			i++
		}
	}

	if len(actions) == 0 {
		var err error
		actions, err = tui.RunWizard()
		if err != nil {
			fmt.Println("Error in TUI:", err)
			os.Exit(1)
		}
	}

	ffmpeg.RunFFmpeg(input, output, actions)
}
