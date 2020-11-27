package dict

import (
	"github.com/psyllos/rally-cli/pkg/context"
	"github.com/spf13/cobra"
)

// NewDictCmd creates a `dict` command
func NewDictCmd(cmdContext *context.CmdContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dict",
		Short: "Manage the data dictionary",
	}

	cmd.AddCommand(NewBuildCmd(cmdContext))
	cmd.AddCommand(NewStatusCmd(cmdContext))

	return cmd
}
