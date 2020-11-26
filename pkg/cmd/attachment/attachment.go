package attachment

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewAttachmentCmd creates a aliasCmd
func NewAttachmentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "attachment",
		Short: "Manage attachments",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("attachment called")
		},
	}

	return cmd
}
