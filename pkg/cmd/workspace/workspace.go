package workspace

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewWorkspaceCmd creates a aliasCmd
func NewWorkspaceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "workspace",
		Short: "Manage workspaces",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("workspace called")
		},
	}

	return cmd
}
