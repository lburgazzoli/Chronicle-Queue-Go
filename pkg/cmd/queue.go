package cmd

import (
	"github.com/spf13/cobra"
)

func NewQueueCommand(root *cobra.Command) *cobra.Command {
	cmd := cobra.Command{
		Use:   "queue",
		Short: "queue",
	}

	cmd.AddCommand(NewQueueTailCommand(root))

	return &cmd
}
