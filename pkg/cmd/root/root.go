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
	"github.com/psyllos/rally-cli/pkg/cmd/subscription"
	"github.com/psyllos/rally-cli/pkg/cmd/task"
	"github.com/psyllos/rally-cli/pkg/cmd/testcase"
	"github.com/psyllos/rally-cli/pkg/cmd/user"
	"github.com/psyllos/rally-cli/pkg/cmd/workspace"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "rally <command> <subcommands> [flags]",
		Short: "Rally CLI",
		Long:  `Rally is a CLI helping you to work seamlessly with Rally from the command line.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		//Run: func(cmd *cobra.Command, args []string) {},
	}

	rootCmd.AddCommand(alias.NewAliasCmd())
	rootCmd.AddCommand(attachment.NewAttachmentCmd())
	rootCmd.AddCommand(config.NewConfigCmd())
	rootCmd.AddCommand(defect.NewDefectCmd())
	rootCmd.AddCommand(defectsuite.NewDefectSuiteCmd())
	rootCmd.AddCommand(hierarchicalrequirement.NewHierarchicalRequirementCmd())
	rootCmd.AddCommand(iteration.NewIterationCmd())
	rootCmd.AddCommand(project.NewProjectCmd())
	rootCmd.AddCommand(release.NewReleaseCmd())
	rootCmd.AddCommand(subscription.NewSubscriptionCmd())
	rootCmd.AddCommand(task.NewTaskCmd())
	rootCmd.AddCommand(testcase.NewTestCaseCmd())
	rootCmd.AddCommand(user.NewUserCmd())
	rootCmd.AddCommand(workspace.NewWorkspaceCmd())

	return rootCmd
}
