package cmd

import (
	"github.com/spf13/cobra"
)

var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "Get information, update or configure the Muthur Command cli backend",
	Long: `
The cli command allows you to manage the internal Muthur Command CLI backend by
exposing commands to view, monitor, configure and control it.`,
	Example: `
  mc cli info
  mc cli update`,
}

func init() {
	rootCmd.AddCommand(cliCmd)
}
