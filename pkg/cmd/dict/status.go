package dict

import (
	"fmt"

	"github.com/psyllos/rally-cli/pkg/context"
	"github.com/spf13/cobra"
)

// NewStatusCmd creates a `status` command
func NewStatusCmd(cmdContext *context.CmdContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Get status of the data dictionary - detects if a new version is available",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("status called")
		},
	}

	return cmd
}
