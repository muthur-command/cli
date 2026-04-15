package cmd

import (
	"log/slog"

	helper "github.com/muthur-command/cli/client"
	"github.com/spf13/cobra"
)

var hostOptionsCmd = &cobra.Command{
	Use:     "options",
	Aliases: []string{"option", "opt", "opts", "op"},
	Short:   "Allow to set options on host system",
	Long: `
This command allows you to set configuration options on the host system that 
your Muthur Command is running on.`,
	Example: `
  mc host options --hostname muthurcommand.local`,
	ValidArgsFunction: cobra.NoFileCompletions,
	Args:              cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Debug("host options", "args", args)

		section := "host"
		command := "options"

		var options map[string]any

		hostname, _ := cmd.Flags().GetString("hostname")
		if hostname != "" {
			options = map[string]any{"hostname": hostname}
		}

		resp, err := helper.GenericJSONPost(section, command, options)
		if err != nil {
			helper.PrintError(err)
			ExitWithError = true
		} else {
			ExitWithError = !helper.ShowJSONResponse(resp)
		}
	},
}

func init() {
	hostOptionsCmd.Flags().StringP("hostname", "", "", "Hostname to set")
	hostOptionsCmd.RegisterFlagCompletionFunc("hostname", cobra.NoFileCompletions)
	hostCmd.AddCommand(hostOptionsCmd)
}
