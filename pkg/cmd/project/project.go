package project

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewProjectCmd creates a aliasCmd
func NewProjectCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "project",
		Short: "Manage projects",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("project called")
		},
	}

	return cmd
}
