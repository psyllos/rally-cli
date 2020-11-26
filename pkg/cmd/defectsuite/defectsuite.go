package defectsuite

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewDefectSuiteCmd creates a aliasCmd
func NewDefectSuiteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "defect-suite",
		Short: "Manage defect suites",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("defect-suite called")
		},
	}

	return cmd
}
