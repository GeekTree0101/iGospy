package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "igospy",
		Short: "igospy is spy object generator for clean swift",
	}
)

// Execute starts the command for iGospy
func Execute() error {
	return rootCmd.Execute()
}
