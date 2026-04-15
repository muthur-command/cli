package cmd

import (
	"log/slog"

	helper "github.com/muthur-command/cli/client"
	"github.com/spf13/cobra"
)

var audioReloadCmd = &cobra.Command{
	Use:     "reload",
	Aliases: []string{"refresh", "re"},
	Short:   "Reload the Muthur Command Audio updating information",
	Long: `
Reloading the Muthur Command Audio, triggers the to regather
all data and devices it currently has.`,
	Example: `
  mc audio reload`,
	ValidArgsFunction: cobra.NoFileCompletions,
	Args:              cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Debug("audio reload", "args", args)

		section := "audio"
		command := "reload"

		resp, err := helper.GenericJSONPost(section, command, nil)
		if err != nil {
			helper.PrintError(err)
			ExitWithError = true
		} else {
			ExitWithError = !helper.ShowJSONResponse(resp)
		}
	},
}

func init() {
	audioCmd.AddCommand(audioReloadCmd)
}
