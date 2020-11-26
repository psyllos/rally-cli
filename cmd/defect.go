package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// defectCmd represents the defect command
var defectCmd = &cobra.Command{
	Use:   "defect",
	Short: "Manage defects",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("defect called")
	},
}

func init() {
	rootCmd.AddCommand(defectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// defectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// defectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
