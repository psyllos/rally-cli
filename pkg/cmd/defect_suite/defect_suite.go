package defectsuite

import (
	"fmt"

	"github.com/psyllos/rally-cli/pkg/context"
	"github.com/spf13/cobra"
)

// NewDefectSuiteCmd creates a `defect-suite` command
func NewDefectSuiteCmd(cmdContext *context.CmdContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "defect-suite",
		Short: "Manage defect suites",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("defect-suite called")
		},
	}

	return cmd
}
