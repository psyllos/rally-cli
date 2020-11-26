package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// hierarchicalRequirementCmd represents the hierarchicalRequirement command
var hierarchicalRequirementCmd = &cobra.Command{
	Use:   "hierarchical-requirement",
	Short: "Manage hierarchical requirements",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hierarchicalRequirement called")
	},
}

func init() {
	rootCmd.AddCommand(hierarchicalRequirementCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hierarchicalRequirementCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hierarchicalRequirementCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
