package root

import (
	"github.com/MakeNowJust/heredoc"
	versionCmd "github.com/psyllos/rally-cli/pkg/cmd/version"
	"github.com/psyllos/rally-cli/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdRoot(f *cmdutil.Factory, version, buildDate string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rally <command> <subcommand> [flags]",
		Short: "Rally CLI",
		Long:  `Work seamlessly with Rally from the command line.`,

		SilenceErrors: true,
		SilenceUsage:  true,
		Example: heredoc.Doc(`
			$ gh issue create
			$ gh repo clone cli/cli
			$ gh pr checkout 321
		`),
		Annotations: map[string]string{
			"help:feedback": heredoc.Doc(`
				Open an issue using 'gh issue create -R github.com/psyllos/rally-cli'
			`),
			"help:environment": heredoc.Doc(`
				See 'rally help environment' for the list of supported environment variables.
			`),
		},
	}

	cmd.SetOut(f.IOStreams.Out)
	cmd.SetErr(f.IOStreams.ErrOut)

	cs := f.IOStreams.ColorScheme()

	helpHelper := func(command *cobra.Command, args []string) {
		rootHelpFunc(cs, command, args)
	}

	cmd.PersistentFlags().Bool("help", false, "Show help for command")
	cmd.SetHelpFunc(helpHelper)
	cmd.SetUsageFunc(rootUsageFunc)
	cmd.SetFlagErrorFunc(rootFlagErrorFunc)

	formattedVersion := versionCmd.Format(version, buildDate)
	cmd.SetVersionTemplate(formattedVersion)
	cmd.Version = formattedVersion
	cmd.Flags().Bool("version", false, "Show gh version")

	// Help topics
	cmd.AddCommand(NewHelpTopic("environment"))
	referenceCmd := NewHelpTopic("reference")
	referenceCmd.SetHelpFunc(referenceHelpFn(f.IOStreams))
	cmd.AddCommand(referenceCmd)

	// this needs to appear last:
	referenceCmd.Long = referenceLong(cmd)
	return cmd
}
