package cmd

import (
	"log/slog"

	helper "github.com/muthur-command/cli/client"
	"github.com/spf13/cobra"
)

var multicastStatsCmd = &cobra.Command{
	Use:     "stats",
	Aliases: []string{"status", "stat"},
	Short:   "Provides system usage stats of the Muthur Command Multicast server",
	Long: `
Provides insight into the system usage stats of the Muthur Command Multicast server.
It shows you how much CPU, memory, disk & network resources it uses.
`,
	Example: `
  ha multicast stats
`,
	ValidArgsFunction: cobra.NoFileCompletions,
	Args:              cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Debug("multicast stats", "args", args)

		section := "multicast"
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
	multicastCmd.AddCommand(multicastStatsCmd)
}
