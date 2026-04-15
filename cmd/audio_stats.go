package cmd

import (
	"log/slog"

	helper "github.com/muthur-command/cli/client"
	"github.com/spf13/cobra"
)

var audioStatsCmd = &cobra.Command{
	Use:     "stats",
	Aliases: []string{"status", "stat", "st"},
	Short:   "Provides system usage stats of Muthur Command Audio",
	Long: `
Provides insight into the system usage stats of Muthur Command Audio.
It shows you how much CPU, memory, disk & network resources it uses.`,
	Example: `
  mc audio stats`,
	ValidArgsFunction: cobra.NoFileCompletions,
	Args:              cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Debug("audio stats", "args", args)

		section := "audio"
		command := "stats"

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
	audioCmd.AddCommand(audioStatsCmd)
}
