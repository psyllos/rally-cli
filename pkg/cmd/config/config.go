package config

import (
	"github.com/psyllos/rally-cli/pkg/context"
	"github.com/spf13/cobra"
)

// NewConfigCmd creates a `config` command
func NewConfigCmd(cmdContext *context.CmdContext) *cobra.Command {
	configCmd := &cobra.Command{
		Use:   "config",
		Short: "Manage configuration for Rally CLI",
	}

	configCmd.AddCommand(NewGetCmd(cmdContext))
	configCmd.AddCommand(NewSetCmd(cmdContext))

	return configCmd
}
