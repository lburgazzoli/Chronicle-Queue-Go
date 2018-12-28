package cmd

import (
	"github.com/spf13/cobra"
)

func NewQueueTailCommand(queue QueueCommand) *cobra.Command {
	impl := QueueTailCommand{
		queue: queue,
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
	queue QueueCommand
}

func (c *QueueTailCommand) Execute(cmd *cobra.Command, args []string) {
}
