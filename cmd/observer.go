package cmd

import (
	"github.com/spf13/cobra"
)

var observerCmd = &cobra.Command{
	Use:   "observer",
	Short: "Get information, update or configure the Muthur Command observer",
	Long: `
The observer command allows you to manage the internal Muthur Command observer by
exposing commands to view, monitor, configure and control it.`,
	Example: `
  mc observer info
  mc observer update`,
}

func init() {
	rootCmd.AddCommand(observerCmd)
}
