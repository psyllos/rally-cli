package alias

import (
	"fmt"

	"github.com/psyllos/rally-cli/pkg/context"
	"github.com/spf13/cobra"
)

// NewAliasCmd creates a `alias` command
func NewAliasCmd(cmdContext *context.CmdContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "alias",
		Short: "Create query shortcuts",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("alias called")
		},
	}

	return cmd
}
