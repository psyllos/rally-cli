package status

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/psyllos/rally-cli/pkg/context"
	"github.com/spf13/cobra"
)

// NewStatusCmd creates a `alias` command
func NewStatusCmd(cmdContext *context.CmdContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Get the status",
		Run: func(cmd *cobra.Command, args []string) {

			var objmap map[string]json.RawMessage

			if err := cmdContext.Client.Request("rally1.rallydev.com", "GET", "slm/webservice/v2.x/user", nil, &objmap); err != nil {
				fmt.Printf("an error occured: %v\n", err)
				os.Exit(1)
			}
		},
	}

	return cmd
}
