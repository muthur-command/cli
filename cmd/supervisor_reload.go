package cmd

import (
	"log/slog"

	helper "github.com/muthur-command/cli/client"
	"github.com/spf13/cobra"
)

var supervisorReloadCmd = &cobra.Command{
	Use:     "reload",
	Aliases: []string{"refresh", "re"},
	Short:   "Reload the Muthur Command Supervisor updating information",
	Long: `
Reloading the Muthur Command Supervisor, triggers the Supervisor to regather
all data it currently has, including checking for updates.`,
	Example: `
  ha supervisor reload`,
	ValidArgsFunction: cobra.NoFileCompletions,
	Args:              cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Debug("supervisor reload", "args", args)

		section := "supervisor"
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
	supervisorCmd.AddCommand(supervisorReloadCmd)
}
