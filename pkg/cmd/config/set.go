package config

import (
	"fmt"

	"github.com/psyllos/rally-cli/pkg/context"

	"github.com/spf13/cobra"
)

// NewSetCmd creates a `set` command
func NewSetCmd(cmdContext *context.CmdContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set",
		Short: "Print the value of a given configuration key",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("set called")
		},
	}

	return cmd
}
