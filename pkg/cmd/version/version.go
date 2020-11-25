package version

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/psyllos/rally-cli/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdVersion(factory *cmdutil.Factory, version, buildDate string) *cobra.Command {
	cmd := &cobra.Command{
		Use:    "version",
		Hidden: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprint(factory.IOStreams.Out, Format(version, buildDate))
		},
	}

	return cmd
}

func Format(version, buildDate string) string {
	version = strings.TrimPrefix(version, "v")

	if buildDate != "" {
		version = fmt.Sprintf("%s (%s)", version, buildDate)
	}

	return fmt.Sprintf("gh version %s\n%s\n", version, changelogURL(version))
}

func changelogURL(version string) string {
	path := "https://github.com/psyllos/rally-cli"
	r := regexp.MustCompile(`^v?\d+\.\d+\.\d+(-[\w.]+)?$`)
	if !r.MatchString(version) {
		return fmt.Sprintf("%s/releases/latest", path)
	}

	url := fmt.Sprintf("%s/releases/tag/v%s", path, strings.TrimPrefix(version, "v"))
	return url
}
