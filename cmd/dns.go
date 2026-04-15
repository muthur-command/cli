package cmd

import (
	"github.com/spf13/cobra"
)

var dnsCmd = &cobra.Command{
	Use:   "dns",
	Short: "Get information, update or configure the Muthur Command DNS server",
	Long: `
The dns command allows you to manage the internal Muthur Command DNS server by
exposing commands to view, monitor, configure and control it.`,
	Example: `
  ha dns logs
  ha dns info
  ha dns update`,
}

func init() {
	rootCmd.AddCommand(dnsCmd)
}
