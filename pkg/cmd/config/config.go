package config

import (
	"github.com/spf13/cobra"
)

// NewConfigCmd creates a configCmd
func NewConfigCmd() *cobra.Command {
	configCmd := &cobra.Command{
		Use:   "config",
		Short: "Manage configuration for Rally CLI",
	}

	configCmd.AddCommand(NewGetCmd())
	configCmd.AddCommand(NewSetCmd())

	return configCmd
}
