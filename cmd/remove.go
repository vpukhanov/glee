package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vpukhanov/glee/pkg/glee"
)

var removeCmd = &cobra.Command{
	Use:   "remove [files to unignore]",
	Short: "Remove entries from git exclude list",
	Long: `Remove files from the local exclude list, allowing them to be tracked by git again.

You can remove multiple files using glob patterns:
	glee remove filename*.txt
		
To remove a glob pattern itself from the exclude list, escape special characters with a backslash:
	glee remove filename\*.txt`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return glee.RemoveExcludes(args)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
