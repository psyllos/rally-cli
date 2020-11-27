package project

import (
	"fmt"

	"github.com/psyllos/rally-cli/pkg/context"
	"github.com/spf13/cobra"
)

// NewProjectCmd creates a `project` command
func NewProjectCmd(cmdContext *context.CmdContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "project",
		Short: "Manage projects",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("project called")
		},
	}

	return cmd
}
