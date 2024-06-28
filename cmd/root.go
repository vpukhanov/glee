package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	version = "dev" // This can be set during build time
)

var rootCmd = &cobra.Command{
	Use:     "glee",
	Short:   "glee - exclude files from git tracking without adding them to .gitignore",
	Version: version,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
