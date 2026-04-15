package cmd

import (
	"log/slog"

	helper "github.com/muthur-command/cli/client"
	"github.com/spf13/cobra"
)

var observerInfoCmd = &cobra.Command{
	Use:     "info",
	Aliases: []string{"in", "inf"},
	Short:   "Shows information about the internal Muthur Command observer",
	Long: `
Shows information about the internally running Muthur Command observer
`,
	Example: `
  mc observer info
`,
	ValidArgsFunction: cobra.NoFileCompletions,
	Args:              cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Debug("observer info", "args", args)

		section := "observer"
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
	observerCmd.AddCommand(observerInfoCmd)
}
