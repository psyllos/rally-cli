package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// testCaseCmd represents the testCase command
var testCaseCmd = &cobra.Command{
	Use:   "test-case",
	Short: "Manage test cases",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("testCase called")
	},
}

func init() {
	rootCmd.AddCommand(testCaseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCaseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCaseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
