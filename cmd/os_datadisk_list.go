package cmd

import (
	"log/slog"

	helper "github.com/muthur-command/cli/client"
	"github.com/spf13/cobra"
)

var osDataDiskListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"in", "inf", "info", "show"},
	Short:   "Provides information about the running Muthur Command Operating System",
	Long: `
This command provides general information about available Harddisk for using with Muthur Command Operating System.
`,
	Example: `
  mc os datadisk list
`,
	ValidArgsFunction: cobra.NoFileCompletions,
	Args:              cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Debug("os datadisk list", "args", args)

		section := "os"
		command := "datadisk/list"

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
	osDataDiskCmd.AddCommand(osDataDiskListCmd)
}
