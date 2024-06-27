package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "glee",
	Short: "glee - exclude files from git tracking without adding them to .gitignore",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Shucks. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
