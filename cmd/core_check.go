package cmd

import (
	"log/slog"

	helper "github.com/muthur-command/cli/client"
	"github.com/spf13/cobra"
)

var coreCheckCmd = &cobra.Command{
	Use:     "check",
	Aliases: []string{"validate", "chk", "ch"},
	Short:   "Validates your Muthur Command Core configuration",
	Long: `
This commands allows you to check/validate your, currently on disk stored,
Muthur Command Core configuration. This is helpful when you've made changes and
want to make sure the configuration is right, before restarting
Muthur Command Core.`,
	Example: `
  mc core check`,
	ValidArgsFunction: cobra.NoFileCompletions,
	Args:              cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Debug("core check", "args", args)

		section := "core"
		command := "check"

		ProgressSpinner.Start()
		resp, err := helper.GenericJSONPostTimeout(section, command, nil, helper.ContainerOperationTimeout)
		ProgressSpinner.Stop()

		if err != nil {
			helper.PrintError(err)
			ExitWithError = true
		} else {
			ExitWithError = !helper.ShowJSONResponse(resp)
		}
	},
}

func init() {
	coreCmd.AddCommand(coreCheckCmd)
}
