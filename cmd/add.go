package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vpukhanov/glee/pkg/glee"
)

var addCmd = &cobra.Command{
	Use:   "add [files to ignore]",
	Short: "Exclude entries from git tracking",
	Long: `Add files to the local exclude list, removing them from git tracking.

You can exclude multiple files using glob patterns:
	glee add filename*.txt
		
To add a glob pattern itself to the exclude list, escape special characters with a backslash:
	glee add filename\*.txt`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return glee.AddExcludes(args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
