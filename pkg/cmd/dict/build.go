package dict

import (
	"encoding/json"
	"fmt"

	"github.com/psyllos/rally-cli/pkg/context"
	"github.com/spf13/cobra"
)

// NewBuildCmd creates a `build` command
func NewBuildCmd(cmdContext *context.CmdContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build",
		Short: "Build the data dictionary",
		Run: func(cmd *cobra.Command, args []string) {
			var objmap map[string]json.RawMessage
			if err := cmdContext.Client.GetCurrentWorkspace(&objmap); err != nil {
				panic(err)
			}

			fmt.Printf("%v", objmap["QueryResult"])
		},
	}

	return cmd
}
