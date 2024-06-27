package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vpukhanov/glee/pkg/glee"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Empties the excluded list",
	RunE: func(cmd *cobra.Command, args []string) error {
		return glee.ClearExcludes()
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}
