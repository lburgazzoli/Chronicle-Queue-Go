package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Execute --
func Execute() {
	chronicle := cobra.Command{
		Use:   "chronicle",
		Short: "chronicle",
	}

	chronicle.AddCommand(NewCompletionCommand(&chronicle))
	chronicle.AddCommand(NewQueueCommand(&chronicle))

	if err := chronicle.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
