package cmd

import (
	"log/slog"

	helper "github.com/muthur-command/cli/client"
	"github.com/spf13/cobra"
)

var dnsResetCmd = &cobra.Command{
	Use:               "reset",
	Short:             "Resets the internal Muthur Command DNS server configuration",
	Long:              `Resets the internal Muthur Command DNS server configuration.`,
	Example:           `mc dns reset`,
	ValidArgsFunction: cobra.NoFileCompletions,
	Args:              cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Debug("dns reset", "args", args)

		section := "dns"
		command := "reset"

		ProgressSpinner.Start()
		resp, err := helper.GenericJSONPost(section, command, nil)
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
	dnsCmd.AddCommand(dnsResetCmd)
}
