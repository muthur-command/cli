package cmd

import (
	"fmt"
	"log/slog"

	"github.com/go-resty/resty/v2"
	"github.com/muthur-command/cli/client"
	"github.com/spf13/cobra"
)

var hostCmd = &cobra.Command{
	Use:     "host",
	Aliases: []string{"ho"},
	Short:   "Control the host/system that Muthur Command is running on",
	Long: `
The host command provides commandline tools to control the host (system) that
Muthur Command is running on. It allows you do thing like reboot or shutdown the
system, but also provides option to change the hostname of the system.`,
	Example: `
  mc host reboot
  mc host options --hostname "muthurcommand.local"
`,
}

func init() {
	rootCmd.AddCommand(hostCmd)
}

func addLogsFlags(cmd *cobra.Command) {
	cmd.Flags().BoolP("follow", "f", false, "Continuously print new log entries")
	cmd.Flags().Uint32P("lines", "n", 0, "Number of log entries to show")
	cmd.Flags().StringP("boot", "b", "", "Logs of particular boot ID")
	cmd.Flags().BoolP("verbose", "v", false, "Return logs in verbose format")
	cmd.Flags().Lookup("follow").NoOptDefVal = "true"
	cmd.Flags().Lookup("verbose").NoOptDefVal = "true"

	cmd.RegisterFlagCompletionFunc("follow", boolCompletions)
	cmd.RegisterFlagCompletionFunc("verbose", boolCompletions)
	cmd.RegisterFlagCompletionFunc("lines", cobra.NoFileCompletions)
	cmd.RegisterFlagCompletionFunc("boot", hostBootCompletions)
}

func processLogsFlags(section string, cmd *cobra.Command) (*resty.Request, error) {
	command := "logs"

	boot, _ := cmd.Flags().GetString("boot")
	if len(boot) > 0 {
		command += "/boots/{boot}"
	}

	follow, _ := cmd.Flags().GetBool("follow")
	if follow {
		command += "/follow"
	}

	URL, err := client.URLHelper(section, command)
	if err != nil {
		return nil, err
	}

	accept := "text/plain"
	verbose, _ := cmd.Flags().GetBool("verbose")
	if verbose {
		accept = "text/x-log"
	}

	/* Disable timeouts to allow following forever */
	request := client.GetRequestTimeout(0).SetHeader("Accept", accept).SetDoNotParseResponse(true)

	lines, _ := cmd.Flags().GetUint32("lines")
	if lines > 0 {
		rangeHeader := fmt.Sprintf("entries=:%d:", -(int(lines) - 1))
		slog.Debug("Range header", "value", rangeHeader)
		request.SetHeader("Range", rangeHeader)
	}

	request.SetPathParam("boot", boot)
	request.URL = URL

	return request, nil
}
