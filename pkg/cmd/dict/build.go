package dict

import (
	"fmt"

	"github.com/psyllos/rally-cli/pkg/context"
	"github.com/psyllos/rally-cli/pkg/model"
	"github.com/spf13/cobra"
)

// NewBuildCmd creates a `build` command
func NewBuildCmd(cmdContext *context.CmdContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build",
		Short: "Build the data dictionary",
		Run: func(cmd *cobra.Command, args []string) {

			type res struct {
				QueryResult model.QueryResult
			}

			var obj res
			if err := cmdContext.Client.GetCurrentWorkspace(&obj); err != nil {
				panic(err)
			}

			fmt.Printf("%v", obj.QueryResult.Results[0]["_refObjectName"])
		},
	}

	return cmd
}
