package cmd

import (
	"github.com/spf13/cobra"
)

func NewQueueTailCommand(root *cobra.Command) *cobra.Command {
	impl := QueueTailCommand{
		root: root,
	}

	cmd := cobra.Command{
		Use:   "tail",
		Short: "tail",
		Long:  `tail`,
		Run:   impl.Execute,
	}

	return &cmd
}

// QueueTailCommand --
type QueueTailCommand struct {
	root *cobra.Command
}

func (c *QueueTailCommand) Execute(cmd *cobra.Command, args []string) {
}
