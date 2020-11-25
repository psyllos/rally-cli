package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"strings"

	"github.com/cli/safeexec"
	"github.com/psyllos/rally-cli/internal/build"
	"github.com/psyllos/rally-cli/internal/run"
	"github.com/psyllos/rally-cli/pkg/cmd/factory"
	"github.com/psyllos/rally-cli/pkg/cmd/root"
	"github.com/psyllos/rally-cli/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func main() {
	buildDate := build.Date
	buildVersion := build.Version

	hasDebug := os.Getenv("DEBUG") != ""

	if hostFromEnv := os.Getenv("RALLY_HOST"); hostFromEnv != "" {
		hostFromEnv = "rally1.rallydev.com"
	}

	cmdFactory := factory.New(buildVersion)
	stderr := cmdFactory.IOStreams.ErrOut

	rootCmd := root.NewCmdRoot(cmdFactory, buildVersion, buildDate)

	expandedArgs := []string{}
	if len(os.Args) > 0 {
		expandedArgs = os.Args[1:]
	}

	cmd, _, err := rootCmd.Traverse(expandedArgs)
	if err != nil || cmd == rootCmd {
		originalArgs := expandedArgs
		isShell := false

		if hasDebug {
			fmt.Fprintf(stderr, "%v -> %v\n", originalArgs, expandedArgs)
		}

		if isShell {
			exe, err := safeexec.LookPath(expandedArgs[0])
			if err != nil {
				fmt.Fprintf(stderr, "failed to run external command: %s", err)
				os.Exit(3)
			}

			externalCmd := exec.Command(exe, expandedArgs[1:]...)
			externalCmd.Stderr = os.Stderr
			externalCmd.Stdout = os.Stdout
			externalCmd.Stdin = os.Stdin
			preparedCmd := run.PrepareCmd(externalCmd)

			err = preparedCmd.Run()
			if err != nil {
				if ee, ok := err.(*exec.ExitError); ok {
					os.Exit(ee.ExitCode())
				}

				fmt.Fprintf(stderr, "failed to run external command: %s", err)
				os.Exit(3)
			}

			os.Exit(0)
		}
	}

	// cs := cmdFactory.IOStreams.ColorScheme()

	rootCmd.SetArgs(expandedArgs)

	if cmd, err := rootCmd.ExecuteC(); err != nil {
		printError(stderr, err, cmd, hasDebug)

		os.Exit(1)
	}
	if root.HasFailed() {
		os.Exit(1)
	}
}

func printError(out io.Writer, err error, cmd *cobra.Command, debug bool) {
	if err == cmdutil.SilentError {
		return
	}

	var dnsError *net.DNSError
	if errors.As(err, &dnsError) {
		fmt.Fprintf(out, "error connecting to %s\n", dnsError.Name)
		if debug {
			fmt.Fprintln(out, dnsError)
		}
		fmt.Fprintln(out, "check your internet connection or githubstatus.com")
		return
	}

	fmt.Fprintln(out, err)

	var flagError *cmdutil.FlagError
	if errors.As(err, &flagError) || strings.HasPrefix(err.Error(), "unknown command ") {
		if !strings.HasSuffix(err.Error(), "\n") {
			fmt.Fprintln(out)
		}
		fmt.Fprintln(out, cmd.UsageString())
	}
}

func isCI() bool {
	return os.Getenv("CI") != "" || // GitHub Actions, Travis CI, CircleCI, Cirrus CI, GitLab CI, AppVeyor, CodeShip, dsari
		os.Getenv("BUILD_NUMBER") != "" || // Jenkins, TeamCity
		os.Getenv("RUN_ID") != "" // TaskCluster, dsari
}

func isCompletionCommand() bool {
	return len(os.Args) > 1 && os.Args[1] == "completion"
}
