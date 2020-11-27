package user

import (
	"fmt"

	"github.com/psyllos/rally-cli/pkg/context"
	"github.com/spf13/cobra"
)

// NewUserCmd creates a `user` command
func NewUserCmd(cmdContext *context.CmdContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "Manage users",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("user called")
		},
	}

	return cmd
}
