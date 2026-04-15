package cmd

import (
	"log/slog"

	helper "github.com/muthur-command/cli/client"
	"github.com/spf13/cobra"
)

var refreshUpdatesCmd = &cobra.Command{
	Use:     "refresh-updates",
	Aliases: []string{"refresh", "refresh_updates"},
	Short:   "Reload stores and version information",
	Long: `
This command reloads information about app repositories and fetches new version files.
	`,
	Example: `
  mc refresh-updates
	`,
	ValidArgsFunction: cobra.NoFileCompletions,
	Args:              cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Debug("refresh_updates", "args", args)

		section := "refresh_updates"
		command := ""

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
	rootCmd.AddCommand(refreshUpdatesCmd)
}
