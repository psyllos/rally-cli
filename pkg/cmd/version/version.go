package version

import (
	"fmt"
	"strings"

	"github.com/psyllos/rally-cli/pkg/context"

	"github.com/spf13/cobra"
)

func NewCmdVersion(cmdContext *context.CmdContext) *cobra.Command {

	cmd := &cobra.Command{
		Use:    "version",
		Hidden: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf(Format(cmdContext.Version, cmdContext.BuildDate))
		},
	}

	return cmd
}

func Format(version, buildDate string) string {
	version = strings.TrimPrefix(version, "v")

	if buildDate != "" {
		version = fmt.Sprintf("%s (%s)", version, buildDate)
	}

	return fmt.Sprintf("Rally CLI - Version %s\n", version)
}
