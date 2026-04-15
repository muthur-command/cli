package cmd

import (
	"log/slog"

	"github.com/spf13/cobra"
)

var supervisorCmd = &cobra.Command{
	Use:     "supervisor",
	Aliases: []string{"super", "su"},
	Short:   "Monitor, control and configure the Muthur Command Supervisor",
	Long: `
The Muthur Command Supervisor is the heart of the Muthur Command system.
It manages your Muthur Command Core, Operating System, and all the apps.
It even manages itself! This series of command give you control over the
Muthur Command Supervisor.`,
	Example: `
  ha supervisor reload
  ha supervisor update
  ha supervisor logs`,
}

func init() {
	slog.Debug("Init supervisor")
	rootCmd.AddCommand(supervisorCmd)
}
