package cmd

import (
	"log/slog"

	helper "github.com/muthur-command/cli/client"
	"github.com/spf13/cobra"
)

var audioRestartCmd = &cobra.Command{
	Use:     "restart",
	Aliases: []string{"reboot"},
	Short:   "Restarts the internal Muthur Command Audio",
	Long:    `Restart the internal Muthur Command Audio`,
	Example: `
  mc audio restart`,
	ValidArgsFunction: cobra.NoFileCompletions,
	Args:              cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Debug("audio restart", "args", args)

		section := "audio"
		command := "restart"

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
	audioCmd.AddCommand(audioRestartCmd)
}
