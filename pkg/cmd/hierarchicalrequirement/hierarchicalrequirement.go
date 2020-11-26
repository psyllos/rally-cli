package hierarchicalrequirement

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewHierarchicalRequirementCmd creates a aliasCmd
func NewHierarchicalRequirementCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hierarchical-requirement",
		Short: "Manage hierarchical requirements",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("hierarchical-requirement called")
		},
	}

	return cmd
}
