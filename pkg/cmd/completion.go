package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func NewCompletionCommand(root *cobra.Command) *cobra.Command {
	cmd := cobra.Command{
		Use:   "completion",
		Short: "Generates shell completion scripts",
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "bash",
		Short: "Generates bash completion scripts",
		Run: func(cmd *cobra.Command, args []string) {
			root.GenBashCompletion(os.Stdout)
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use:   "zsh",
		Short: "Generates zsh completion scripts",
		Run: func(cmd *cobra.Command, args []string) {
			root.GenZshCompletion(os.Stdout)
		},
	})

	return &cmd
}
