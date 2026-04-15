package cmd

import (
	"log/slog"

	helper "github.com/muthur-command/cli/client"
	"github.com/spf13/cobra"
)

var osUpdateCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"upgrade", "downgrade", "up", "down"},
	Short:   "Updates the Muthur Command Operating System",
	Long: `
Using this command you can upgrade or downgrade the Muthur Command 
Operating System to the latest version or the version specified.
`,
	Example: `
  mc os update
  mc os update --version 5
`,
	ValidArgsFunction: cobra.NoFileCompletions,
	Args:              cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Debug("os update", "args", args)

		section := "os"
		command := "update"

		var options map[string]any

		version, _ := cmd.Flags().GetString("version")
		if version != "" {
			options = map[string]any{"version": version}
		}

		ProgressSpinner.Start()
		resp, err := helper.GenericJSONPostTimeout(section, command, options, helper.OsDownloadTimeout)
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
	osUpdateCmd.Flags().StringP("version", "", "", "Version to update to")
	osUpdateCmd.RegisterFlagCompletionFunc("version", cobra.NoFileCompletions)
	osCmd.AddCommand(osUpdateCmd)
}
