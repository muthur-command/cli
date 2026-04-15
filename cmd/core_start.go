package cmd

import (
	"log/slog"

	helper "github.com/muthur-command/cli/client"
	"github.com/spf13/cobra"
)

var coreStartCmd = &cobra.Command{
	Use:     "start",
	Aliases: []string{"run", "st"},
	Short:   "Manually start Muthur Command Core",
	Long: `
This command allows you to manually start the Muthur Command Core instance on
your system. This, of course, only applies when it has been stopped.`,
	Example: `
  mc core start`,
	ValidArgsFunction: cobra.NoFileCompletions,
	Args:              cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Debug("core start", "args", args)

		section := "core"
		command := "start"

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
	coreCmd.AddCommand(coreStartCmd)
}
