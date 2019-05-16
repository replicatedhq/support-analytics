package cmd

import (
	"fmt"
	"github.com/replicatedcom/support-analytics/bundle"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

var AnalyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyze a support bundle",
	Long:  `Take apart the support bundle and report back findings.`,
	Run: func(cmd *cobra.Command, args []string) {
		basePath := cmd.Flag("path").Value.String()

		// Verify the base path exists and is valid
		if basePath == "" {
			fmt.Print("Error: missing required field path\n")
			os.Exit(1)
		}
		_, err := ioutil.ReadDir(basePath)
		if err != nil {
			fmt.Print("Error: path is not valid\n")
			os.Exit(1)
		}
		b := bundle.Analyze{BasePath: basePath}
		b.AnalyzeSupportBundle()
	},
}

func init() {
	AnalyzeCmd.PersistentFlags().StringP("path", "p", "", "path to support bundle")
}
