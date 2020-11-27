package defect

import (
	"fmt"

	"github.com/psyllos/rally-cli/pkg/context"
	"github.com/spf13/cobra"
)

// NewDefectCmd creates a `defect` command
func NewDefectCmd(cmdContext *context.CmdContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "defect",
		Short: "Manage defects",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("defect called")
		},
	}

	return cmd
}
