package cmd

import (
	"log/slog"

	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:     "authentication",
	Aliases: []string{"auth", "au"},
	Short:   "Authentication for Muthur Command users.",
	Long: `
The authentication command allows you to manage Muthur Command user accounts.
`,
	Example: `
  ha authentication reset --username "JohnDoe" --password "123SuperSecret!"
	`,
}

func init() {
	slog.Debug("Init authentication")

	rootCmd.AddCommand(authCmd)
}
