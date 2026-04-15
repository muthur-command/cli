package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var coreCmd = &cobra.Command{
	Use:     "core",
	Aliases: []string{"homeassistant", "home-assistant", "ha"},
	Short:   "Provides control of the Muthur Command Core",
	Long: `
This command provides a set of subcommands to control the Muthur Command Core
instance running on this installation.

It provides commands to control Muthur Command Core (start, stop, restart),
but also allows you to check your Muthur Command Core configuration.
Furthermore, some options can be set and allows for upgrading/downgrading
Muthur Command Core.
`,
	Example: `
  ha core check
  ha core restart
  ha core update
	ha core update --version 2021.11.5`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		for idx, arg := range os.Args {
			if idx != 0 && (arg == "homeassistant" || arg == "ha") {
				cmd.PrintErrf("The use of '%s' is deprecated, please use 'core' instead!\n", arg)
			}
		}
		rootCmd.PersistentPreRun(cmd, args)
	},
}

func init() {
	// add cmd to root command
	rootCmd.AddCommand(coreCmd)
}
