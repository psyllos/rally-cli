package task

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewTaskCmd creates a aliasCmd
func NewTaskCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "task",
		Short: "Manage tasks",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("task called")
		},
	}

	return cmd
}
