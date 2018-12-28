package cmd

import (
	"github.com/spf13/cobra"
)

func NewQueueCommand(root *cobra.Command) *cobra.Command {
	cmd := cobra.Command{
		Use:   "queue",
		Short: "queue",
	}

	impl := QueueCommand{
		Root: root,
	}

	//
	// TODO: add persistent flags
	//

	// add sub commands
	cmd.AddCommand(NewQueueInfoCommand(impl))
	cmd.AddCommand(NewQueueTailCommand(impl))

	return &cmd
}

// QueueCommand --
type QueueCommand struct {
	Root *cobra.Command
}
