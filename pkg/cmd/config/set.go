package config

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewSetCmd creates a `set` command
func NewSetCmd() *cobra.Command {
	setCmd := &cobra.Command{
		Use:   "set",
		Short: "Print the value of a given configuration key",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("set called")
		},
	}

	return setCmd
}
