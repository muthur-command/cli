package cmd

import (
	"log/slog"

	helper "github.com/muthur-command/cli/client"
	"github.com/spf13/cobra"
)

var coreInfoCmd = &cobra.Command{
	Use:     "info",
	Aliases: []string{"in", "inf"},
	Short:   "Provides information about Muthur Command Core",
	Long: `
This command provides information about the running Muthur Command Core instance
running on your Muthur Command system.`,
	Example: `
  mc core info`,
	ValidArgsFunction: cobra.NoFileCompletions,
	Args:              cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Debug("core info", "args", args)

		section := "core"
		command := "info"

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
	coreCmd.AddCommand(coreInfoCmd)
}
