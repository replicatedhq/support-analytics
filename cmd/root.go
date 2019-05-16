package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "support",
	Short: "Support analysis tool",
	Long:  `This tool is for taking apart support bundles`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}
