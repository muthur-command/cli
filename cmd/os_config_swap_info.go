package cmd

import (
	"log/slog"

	helper "github.com/muthur-command/cli/client"
	"github.com/spf13/cobra"
)

var osConfigSwapInfoCmd = &cobra.Command{
	Use:     "info",
	Aliases: []string{"in", "info"},
	Short:   "Show Muthur Command OS swap settings",
	Long: `
This command allows you to see how swap is used by the Muthur Command OS.`,
	Example: `
  mc os config swap info`,
	ValidArgsFunction: cobra.NoFileCompletions,
	Args:              cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Debug("os config swap info", "args", args)

		section := "os"
		command := "config/swap"

		resp, err := helper.GenericJSONGet(section, command)
		if err != nil {
			helper.PrintError(err)
			ExitWithError = true
		} else {
			ExitWithError = !helper.ShowJSONResponse(resp)
		}
	},
}

func init() {
	osConfigSwapCmd.AddCommand(osConfigSwapInfoCmd)
}
