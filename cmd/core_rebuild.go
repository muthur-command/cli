package cmd

import (
	"log/slog"

	helper "github.com/muthur-command/cli/client"
	"github.com/spf13/cobra"
)

var coreRebuildCmd = &cobra.Command{
	Use:     "rebuild",
	Aliases: []string{"rb", "reinstall"},
	Short:   "Rebuild the Muthur Command Core instance",
	Long: `
This command allows you to trigger a rebuild for your Muthur Command Core
instance running on your Muthur Command system.
Don't worry, this does not delete your config.`,
	Example: `
  mc core rebuild`,
	ValidArgsFunction: cobra.NoFileCompletions,
	Args:              cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Debug("core rebuild", "args", args)

		section := "core"
		command := "rebuild"

		options := make(map[string]any)

		safeMode, err := cmd.Flags().GetBool("safe-mode")
		if err == nil && safeMode {
			options["safe_mode"] = safeMode
		}
		force, err := cmd.Flags().GetBool("force")
		if err == nil && force {
			options["force"] = force
		}

		ProgressSpinner.Start()
		resp, err := helper.GenericJSONPostTimeout(section, command, options, helper.ContainerOperationTimeout)
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
	coreRebuildCmd.Flags().BoolP("safe-mode", "s", false, "Rebuild Muthur Command in safe mode")
	coreRebuildCmd.Flags().BoolP("force", "f", false, "Force rebuild during an offline db migration")
	coreRebuildCmd.Flags().Lookup("safe-mode").NoOptDefVal = "true"
	coreRebuildCmd.Flags().Lookup("force").NoOptDefVal = "true"
	coreRebuildCmd.RegisterFlagCompletionFunc("safe-mode", boolCompletions)
	coreRebuildCmd.RegisterFlagCompletionFunc("force", boolCompletions)

	coreCmd.AddCommand(coreRebuildCmd)
}
