package task

import (
	"fmt"

	"github.com/psyllos/rally-cli/pkg/context"
	"github.com/spf13/cobra"
)

// NewTaskCmd creates a `task` command
func NewTaskCmd(cmdContext *context.CmdContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "task",
		Short: "Manage tasks",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("task called")
		},
	}

	return cmd
}
