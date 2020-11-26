package alias

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewAliasCmd creates a aliasCmd
func NewAliasCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "alias",
		Short: "Create query shortcuts",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("alias called")
		},
	}

	return cmd
}
