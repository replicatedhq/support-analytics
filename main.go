package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/replicatedcom/support-analytics/cmd"
)

func main() {
	cmd.RootCmd.PersistentFlags().Bool("debug", false, "debug logging")

	// Set the log output via the global flag
	cmd.RootCmd.ParseFlags(os.Args)
	if cmd.RootCmd.Flag("debug").Value.String() == "true" {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}

	// Add available commands
	cmd.RootCmd.AddCommand(cmd.AnalyzeCmd)
	cmd.RootCmd.AddCommand(cmd.VersionCmd)
	cmd.RootCmd.AddCommand(cmd.EventsCmd)
	cmd.RootCmd.AddCommand(cmd.AppYamlCommand)

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
