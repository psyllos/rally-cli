package workspace

import (
	"fmt"

	"github.com/psyllos/rally-cli/pkg/context"
	"github.com/spf13/cobra"
)

// NewWorkspaceCmd creates a `workspace` command
func NewWorkspaceCmd(cmdContext *context.CmdContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "workspace",
		Short: "Manage workspaces",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("workspace called")
		},
	}

	return cmd
}
