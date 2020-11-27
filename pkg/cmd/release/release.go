package release

import (
	"fmt"

	"github.com/psyllos/rally-cli/pkg/context"
	"github.com/spf13/cobra"
)

// NewReleaseCmd creates a `release` command
func NewReleaseCmd(cmdContext *context.CmdContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "release",
		Short: "Manage releases",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("release called")
		},
	}

	return cmd
}
