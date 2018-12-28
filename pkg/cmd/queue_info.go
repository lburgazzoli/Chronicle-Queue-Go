package cmd

import (
	"github.com/spf13/cobra"
)

func NewQueueInfoCommand(queue QueueCommand) *cobra.Command {
	impl := QueueInfoCommand{
		queue: queue,
	}

	cmd := cobra.Command{
		Use:   "info",
		Short: "info",
		Long:  `info`,
		Run:   impl.Execute,
	}

	return &cmd
}

// QueueInfoCommand --
type QueueInfoCommand struct {
	queue QueueCommand
}

func (c *QueueInfoCommand) Execute(cmd *cobra.Command, args []string) {
}
