package cmd

import (
	"log/slog"

	helper "github.com/muthur-command/cli/client"
	"github.com/spf13/cobra"
)

var backupsInfoCmd = &cobra.Command{
	Use:     "info [slug]",
	Aliases: []string{"in", "inf"},
	Short:   "Provides information about the current available backups",
	Long: `
When a Muthur Command backup is created, it will be available for restore.
This command gives you information about a specific backup.`,
	Example: `
  mc backups info c1a07617`,
	ValidArgsFunction: backupsCompletions,
	Args:              cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		slog.Debug("backups info", "args", args)

		section := "backups"
		command := "{slug}/info"

		url, err := helper.URLHelper(section, command)

		if err != nil {
			helper.PrintError(err)
			ExitWithError = true
			return
		}

		request := helper.GetJSONRequest()

		slug := args[0]

		request.SetPathParams(map[string]string{
			"slug": slug,
		})

		resp, err := request.Get(url)
		resp, err = helper.GenericJSONErrorHandling(resp, err)

		if err != nil {
			helper.PrintError(err)
			ExitWithError = true
		} else {
			ExitWithError = !helper.ShowJSONResponse(resp)
		}
	},
}

func init() {
	backupsCmd.AddCommand(backupsInfoCmd)
}
