package cmd

import (
	"github.com/spf13/cobra"
)

var hostDisksCmd = &cobra.Command{
	Use:     "disks",
	Aliases: []string{"disk"},
	Short:   "Manage host disk operations",
	Long: `
The disks command provides access to disk-related operations on the host system
that Muthur Command is running on.`,
	Example: `
  mc host disks usage`,
}

func init() {
	hostCmd.AddCommand(hostDisksCmd)
}
