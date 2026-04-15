package cmd

import (
	"log/slog"

	helper "github.com/muthur-command/cli/client"
	"github.com/spf13/cobra"
)

var securityInfoCmd = &cobra.Command{
	Use:     "info",
	Aliases: []string{"in", "inf"},
	Short:   "Provides information about the Muthur Command Security backend",
	Long: `
This command provides general information about the Muthur Command Security backend.
`,
	Example: `
  mc security info
`,
	ValidArgsFunction: cobra.NoFileCompletions,
	Args:              cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Debug("security info", "args", args)

		section := "security"
		command := "info"

		resp, err := helper.GenericJSONGet(section, command)
		if err != nil {
			helper.PrintError(err)
			ExitWithError = true
		} else {
			ExitWithError = !helper.ShowJSONResponse(resp)
		}
	},
}

func init() {
	securityCmd.AddCommand(securityInfoCmd)
}
