package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// defectSuiteCmd represents the defectSuite command
var defectSuiteCmd = &cobra.Command{
	Use:   "defect-suite",
	Short: "Manage defect suites",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("defectSuite called")
	},
}

func init() {
	rootCmd.AddCommand(defectSuiteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// defectSuiteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// defectSuiteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
