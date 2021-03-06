package testcase

import (
	"fmt"

	"github.com/psyllos/rally-cli/pkg/context"
	"github.com/spf13/cobra"
)

// NewTestCaseCmd creates a `test-case` command
func NewTestCaseCmd(cmdContext *context.CmdContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "test-case",
		Short: "Manage test cases",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("test-case called")
		},
	}

	return cmd
}
