package cmd

import (
	"github.com/spf13/cobra"
)

var osConfigSwapCmd = &cobra.Command{
	Use:     "swap",
	Aliases: []string{"sw"},
	Short:   "Show or change Muthur Command OS swap settings",
	Long: `
This command allows you to show or change current swap configuration
of Muthur Command OS.`,
	Example: `
  mc os config swap info`,
}

func init() {
	osConfigCmd.AddCommand(osConfigSwapCmd)
}
