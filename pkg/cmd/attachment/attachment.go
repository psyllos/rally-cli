package attachment

import (
	"fmt"

	"github.com/psyllos/rally-cli/pkg/context"
	"github.com/spf13/cobra"
)

// NewAttachmentCmd creates a `attachment` command
func NewAttachmentCmd(cmdContext *context.CmdContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "attachment",
		Short: "Manage attachments",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("attachment called")
		},
	}

	return cmd
}
