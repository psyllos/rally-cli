package root

import (
	"github.com/psyllos/rally-cli/pkg/cmd/alias"
	"github.com/psyllos/rally-cli/pkg/cmd/attachment"
	"github.com/psyllos/rally-cli/pkg/cmd/config"
	"github.com/psyllos/rally-cli/pkg/cmd/defect"
	"github.com/psyllos/rally-cli/pkg/cmd/defectsuite"
	"github.com/psyllos/rally-cli/pkg/cmd/hierarchicalrequirement"
	"github.com/psyllos/rally-cli/pkg/cmd/iteration"
	"github.com/psyllos/rally-cli/pkg/cmd/project"
	"github.com/psyllos/rally-cli/pkg/cmd/release"
	"github.com/psyllos/rally-cli/pkg/cmd/status"
	"github.com/psyllos/rally-cli/pkg/cmd/subscription"
	"github.com/psyllos/rally-cli/pkg/cmd/task"
	"github.com/psyllos/rally-cli/pkg/cmd/testcase"
	"github.com/psyllos/rally-cli/pkg/cmd/user"
	"github.com/psyllos/rally-cli/pkg/cmd/version"
	"github.com/psyllos/rally-cli/pkg/cmd/workspace"
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
	cmd.AddCommand(attachment.NewAttachmentCmd(cmdContext))
	cmd.AddCommand(config.NewConfigCmd(cmdContext))
	cmd.AddCommand(defect.NewDefectCmd(cmdContext))
	cmd.AddCommand(defectsuite.NewDefectSuiteCmd(cmdContext))
	cmd.AddCommand(hierarchicalrequirement.NewHierarchicalRequirementCmd(cmdContext))
	cmd.AddCommand(iteration.NewIterationCmd(cmdContext))
	cmd.AddCommand(project.NewProjectCmd(cmdContext))
	cmd.AddCommand(release.NewReleaseCmd(cmdContext))
	cmd.AddCommand(status.NewStatusCmd(cmdContext))
	cmd.AddCommand(subscription.NewSubscriptionCmd(cmdContext))
	cmd.AddCommand(task.NewTaskCmd(cmdContext))
	cmd.AddCommand(testcase.NewTestCaseCmd(cmdContext))
	cmd.AddCommand(user.NewUserCmd(cmdContext))
	cmd.AddCommand(workspace.NewWorkspaceCmd(cmdContext))

	formattedVersion := version.Format(cmdContext.Version, cmdContext.BuildDate)
	cmd.SetVersionTemplate(formattedVersion)
	cmd.Version = formattedVersion

	cmd.AddCommand(version.NewCmdVersion(cmdContext))

	return cmd
}
