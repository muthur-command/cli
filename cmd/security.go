package cmd

import (
	"github.com/spf13/cobra"
)

var securityCmd = &cobra.Command{
	Use:     "security",
	Aliases: []string{"secure", "sec"},
	Short:   "Get information and manage security functionality",
	Long: `
The security command allows you to manage the internal Muthur Command Security backend and
exposing commands to view, configure and control it.`,
	Example: `
  mc security info
  mc security options`,
}

func init() {
	rootCmd.AddCommand(securityCmd)
}
