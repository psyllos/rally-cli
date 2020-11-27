package root

import (
	"github.com/psyllos/rally-cli/pkg/cmd/alias"
	"github.com/psyllos/rally-cli/pkg/cmd/config"
	"github.com/psyllos/rally-cli/pkg/cmd/dict"
	"github.com/psyllos/rally-cli/pkg/cmd/status"
	"github.com/psyllos/rally-cli/pkg/cmd/version"
	"github.com/psyllos/rally-cli/pkg/context"
	"github.com/spf13/cobra"
)

// NewRootCmd creates a `rally` command
func NewRootCmd(cmdContext *context.CmdContext) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "rally <command> <subcommands> [flags]",
		Short: "Rally CLI",
		Long:  `Rally is a CLI helping you to work seamlessly with CA Rally Software from the command line.`,
	}

	cmd.AddCommand(alias.NewAliasCmd(cmdContext))
	cmd.AddCommand(config.NewConfigCmd(cmdContext))
	cmd.AddCommand(dict.NewDictCmd(cmdContext))
	cmd.AddCommand(status.NewStatusCmd(cmdContext))

	formattedVersion := version.Format(cmdContext.Version, cmdContext.BuildDate)
	cmd.SetVersionTemplate(formattedVersion)
	cmd.Version = formattedVersion

	cmd.AddCommand(version.NewCmdVersion(cmdContext))

	return cmd
}
