package testcase

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewTestCaseCmd creates a aliasCmd
func NewTestCaseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "test-case",
		Short: "Manage test cases",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("test-case called")
		},
	}

	return cmd
}
