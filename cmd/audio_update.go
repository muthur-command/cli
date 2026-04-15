package cmd

import (
	"log/slog"

	helper "github.com/muthur-command/cli/client"
	"github.com/spf13/cobra"
)

var audioUpdateCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"upgrade", "downgrade", "up", "down"},
	Short:   "Updates the Muthur Command Audio",
	Long: `
Using this command you can upgrade or downgrade the Muthur Command Audio
instance running on your system to the latest version or the version specified.`,
	Example: `
  mc audio update
  mc audio update --version 6`,
	ValidArgsFunction: cobra.NoFileCompletions,
	Args:              cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Debug("audio update", "args", args)

		section := "audio"
		command := "update"

		var options map[string]any

		version, _ := cmd.Flags().GetString("version")
		if version != "" {
			options = map[string]any{"version": version}
		}

		ProgressSpinner.Start()
		resp, err := helper.GenericJSONPostTimeout(section, command, options, helper.ContainerDownloadTimeout)
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
	audioUpdateCmd.Flags().StringP("version", "", "", "Version to update to")
	audioUpdateCmd.RegisterFlagCompletionFunc("version", cobra.NoFileCompletions)
	audioCmd.AddCommand(audioUpdateCmd)
}
