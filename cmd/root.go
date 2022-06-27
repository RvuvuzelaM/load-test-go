package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vuvuzela/loadtest/cmd/restapi"
)

var rootCmd = &cobra.Command{
	Use:   "loadtest",
	Short: "Load test allows you to test quickly load test your apps",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(restapi.LoadTestCmd)
}
