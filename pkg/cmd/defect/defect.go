package defect

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewDefectCmd creates a aliasCmd
func NewDefectCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "defect",
		Short: "Manage defects",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("defect called")
		},
	}

	return cmd
}
