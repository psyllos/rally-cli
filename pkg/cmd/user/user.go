package user

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewUserCmd creates a aliasCmd
func NewUserCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "Manage users",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("user called")
		},
	}

	return cmd
}
