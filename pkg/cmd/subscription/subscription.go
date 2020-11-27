package subscription

import (
	"fmt"

	"github.com/psyllos/rally-cli/pkg/context"
	"github.com/spf13/cobra"
)

// NewSubscriptionCmd creates a `subscription` command
func NewSubscriptionCmd(cmdContext *context.CmdContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscription",
		Short: "Manage subscriptions",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("subscription called")
		},
	}

	return cmd
}
