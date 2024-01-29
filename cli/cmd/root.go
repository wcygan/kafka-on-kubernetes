package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "A tool to interact with the Kafka cluster",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(listTopicsCmd)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
