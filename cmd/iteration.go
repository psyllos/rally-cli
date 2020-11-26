package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// iterationCmd represents the iteration command
var iterationCmd = &cobra.Command{
	Use:   "iteration",
	Short: "Manage iterations",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("iteration called")
	},
}

func init() {
	rootCmd.AddCommand(iterationCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// iterationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// iterationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
