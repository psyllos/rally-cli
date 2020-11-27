package hierarchicalrequirement

import (
	"fmt"

	"github.com/psyllos/rally-cli/pkg/context"
	"github.com/spf13/cobra"
)

// NewHierarchicalRequirementCmd creates a `hierarchical-requirement` command
func NewHierarchicalRequirementCmd(cmdContext *context.CmdContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hierarchical-requirement",
		Short: "Manage hierarchical requirements",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("hierarchical-requirement called")
		},
	}

	return cmd
}
