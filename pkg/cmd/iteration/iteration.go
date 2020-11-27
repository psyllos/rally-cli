package iteration

import (
	"fmt"

	"github.com/psyllos/rally-cli/pkg/context"
	"github.com/spf13/cobra"
)

// NewIterationCmd creates a `iteration` command
func NewIterationCmd(cmdContext *context.CmdContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "iteration",
		Short: "Manage iterations",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("iteration called")
		},
	}

	return cmd
}
