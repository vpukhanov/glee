package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vpukhanov/glee/pkg/glee"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List excluded entries",
	RunE: func(cmd *cobra.Command, args []string) error {
		return glee.ListExcludes()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
