package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use:   "cm",
	Short: "",
	Long:  "",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	createCharSystemCommands()
	rootCmd.AddCommand(&dumpCmd)
}
