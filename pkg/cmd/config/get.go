package config

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewGetCmd creates a `get` command
func NewGetCmd() *cobra.Command {
	newGetCmd := &cobra.Command{
		Use:   "get",
		Short: "Print the value of a given configuration key",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("get called")
		},
	}

	return newGetCmd
}
