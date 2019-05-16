package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version number",
	Long:  `Print the version number and exit.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("0.0.1\n")
	},
}
