package cmd

import (
	"github.com/spf13/cobra"
)

var osConfigCmd = &cobra.Command{
	Use:     "config",
	Aliases: []string{"conf", "cfg"},
	Short:   "Show or change Muthur Command OS settings",
	Long: `
This command allows you to show or change settings of Muthur Command OS.`,
	Example: `
  mc os config swap`,
}

func init() {
	osCmd.AddCommand(osConfigCmd)
}
