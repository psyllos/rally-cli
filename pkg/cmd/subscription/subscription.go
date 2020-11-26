package subscription

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewSubscriptionCmd creates a aliasCmd
func NewSubscriptionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscription",
		Short: "Manage subscriptions",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("subscription called")
		},
	}

	return cmd
}
