package iteration

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewIterationCmd creates a aliasCmd
func NewIterationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "iteration",
		Short: "Manage iterations",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("iteration called")
		},
	}

	return cmd
}
