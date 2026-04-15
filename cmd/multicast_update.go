package cmd

import (
	"log/slog"

	helper "github.com/muthur-command/cli/client"
	"github.com/spf13/cobra"
)

var multicastUpdateCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"upgrade", "downgrade", "up", "down"},
	Short:   "Updates the internal Muthur Command Multicast server",
	Long: `
Using this command you can upgrade or downgrade the internal Muthur Command 
Multicast server, to the latest version or the version specified.
`,
	Example: `
  mc multicast update --version 5
`,
	ValidArgsFunction: cobra.NoFileCompletions,
	Args:              cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Debug("multicast update", "args", args)

		section := "multicast"
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
	multicastUpdateCmd.Flags().StringP("version", "", "", "Version to update to")
	multicastUpdateCmd.RegisterFlagCompletionFunc("version", cobra.NoFileCompletions)
	multicastCmd.AddCommand(multicastUpdateCmd)
}
