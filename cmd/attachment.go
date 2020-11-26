package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// attachmentCmd represents the attachment command
var attachmentCmd = &cobra.Command{
	Use:   "attachment",
	Short: "Manage attachments",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("attachment called")
	},
}

func init() {
	rootCmd.AddCommand(attachmentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// attachmentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// attachmentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
