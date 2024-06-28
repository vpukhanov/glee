package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vpukhanov/glee/pkg/glee"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Open exclude file in the text editor",
	RunE: func(cmd *cobra.Command, args []string) error {
		return glee.EditExcludes()
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
