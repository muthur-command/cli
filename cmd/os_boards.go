package cmd

import (
	"github.com/spf13/cobra"
)

var osBoardsCmd = &cobra.Command{
	Use:   "boards",
	Short: "See or change settings of the current board",
	Long: `
This command allows you to see or change settings of the board that Muthur
Command is running on.`,
	Example: `
  mc os boards yellow`,
}

func init() {
	osCmd.AddCommand(osBoardsCmd)
}
