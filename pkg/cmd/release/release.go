package release

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewReleaseCmd creates a aliasCmd
func NewReleaseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "release",
		Short: "Manage releases",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("release called")
		},
	}

	return cmd
}
