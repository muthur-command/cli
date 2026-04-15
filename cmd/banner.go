package cmd

import (
	"fmt"
	"log/slog"
	"strings"
	"time"

	helper "github.com/muthur-command/cli/client"
	"github.com/spf13/cobra"
)

const haBanner = `
███╗   ███╗██╗   ██╗████████╗██╗  ██╗██╗   ██╗██████╗ 
████╗ ████║██║   ██║╚══██╔══╝██║  ██║██║   ██║██╔══██╗
██╔████╔██║██║   ██║   ██║   ███████║██║   ██║██████╔╝
██║╚██╔╝██║██║   ██║   ██║   ██╔══██║██║   ██║██╔══██╗
██║ ╚═╝ ██║╚██████╔╝   ██║   ██║  ██║╚██████╔╝██║  ██║
╚═╝     ╚═╝ ╚═════╝    ╚═╝   ╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═╝

 ██████╗ ██████╗ ███╗   ███╗███╗   ███╗ █████╗ ███╗   ██╗██████╗ 
██╔════╝██╔═══██╗████╗ ████║████╗ ████║██╔══██╗████╗  ██║██╔══██╗
██║     ██║   ██║██╔████╔██║██╔████╔██║███████║██╔██╗ ██║██║  ██║
██║     ██║   ██║██║╚██╔╝██║██║╚██╔╝██║██╔══██║██║╚██╗██║██║  ██║
╚██████╗╚██████╔╝██║ ╚═╝ ██║██║ ╚═╝ ██║██║  ██║██║ ╚████║██████╔╝
 ╚═════╝ ╚═════╝ ╚═╝     ╚═╝╚═╝     ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝╚═════╝

Welcome to the Muthur Command command line interface.
`

func supervisorGet(section string, command string) (outdata *(map[string]any), err error) {
	resp, err := helper.GenericJSONGet(section, command)
	if err != nil {
		return nil, err
	}

	var data *helper.Response
	if resp.IsSuccess() {
		data = resp.Result().(*helper.Response)
	} else {
		data = resp.Error().(*helper.Response)
	}
	if data.Result == "ok" {
		if len(data.Data) > 0 {
			outdata = &(data.Data)
		}
	} else {
		return nil, fmt.Errorf("error returned from Supervisor: %s", data.Message)
	}
	return outdata, nil
}

func getAddresses(addresses []any) string {
	addresses_str := make([]string, len(addresses))
	for i, v := range addresses {
		addresses_str[i] = fmt.Sprint(v)
	}
	return strings.Join(addresses_str, ", ")
}

var bannerCmd = &cobra.Command{
	Use:     "banner",
	Aliases: []string{"ba"},
	Short:   "Prints the CLI Muthur Command banner along with some useful information",
	Example: `
  mc banner
	`,
	ValidArgsFunction: cobra.NoFileCompletions,
	Args:              cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Debug("info", "args", args)

		fmt.Print(haBanner)
		fmt.Println()

		nowait, err := cmd.Flags().GetBool("no-wait")
		if err != nil {
			helper.PrintError(err)
			ExitWithError = true
		}

		var netinfo *(map[string]any)
		if !nowait {
			var supervisorReady bool

			for i := range 180 { // 3 minutes timeout
				var err error
				netinfo, err = supervisorGet("network", "info")
				if err == nil && netinfo != nil {
					netifaces, exist := (*netinfo)["interfaces"]
					if exist && len(netifaces.([]any)) > 0 {
						fmt.Println("Muthur Command Supervisor is running!")
						supervisorReady = true
						break
					}
				}
				if i == 0 {
					fmt.Println("Waiting for Supervisor to start...")
				}
				time.Sleep(1 * time.Second)
			}

			if !supervisorReady {
				fmt.Println("Supervisor is taking longer than expected to start. Use 'mc supervisor logs' to check logs.")
				return
			}
		}

		fmt.Println("System information:")
		if netinfo == nil {
			var err error
			netinfo, err = supervisorGet("network", "info")
			if err != nil {
				fmt.Printf("  Network information unavailable: %s\n", err)
				ExitWithError = true
				return
			}
		}

		// Print network address information
		netifaces, exist := (*netinfo)["interfaces"]
		if exist {
			for _, netiface := range netifaces.([]any) {
				nf := netiface.(map[string]any)
				title_ipv4 := fmt.Sprintf("IPv4 addresses for %s:", nf["interface"])
				title_ipv6 := fmt.Sprintf("IPv6 addresses for %s:", nf["interface"])

				if nf["ipv4"] == nil {
					fmt.Printf("  %-25s (No address)\n", title_ipv4)
				} else {
					ipv4 := nf["ipv4"].(map[string]any)
					ipv4_addresses := ipv4["address"].([]any)
					if len(ipv4_addresses) > 0 {
						fmt.Printf("  %-25s %s\n", title_ipv4, getAddresses(ipv4_addresses))
					} else {
						fmt.Printf("  %-25s (No address)\n", title_ipv4)
					}
				}

				if nf["ipv6"] != nil {
					ipv6 := nf["ipv6"].(map[string]any)
					ipv6_addresses := ipv6["address"].([]any)
					if len(ipv6_addresses) > 0 {
						fmt.Printf("  %-25s %s\n", title_ipv6, getAddresses(ipv6_addresses))
					}
				}
			}
		} else {
			fmt.Printf("  (No networking information)")
		}
		fmt.Println()

		// Print Host URL
		hostinfo, err := supervisorGet("host", "info")
		if err != nil {
			helper.PrintError(err)
			ExitWithError = true
			return
		}
		if hostinfo == nil {
			return
		}
		coreinfo, err := supervisorGet("core", "info")
		if err != nil {
			helper.PrintError(err)
			ExitWithError = true
			return
		}
		if coreinfo == nil {
			return
		}

		protocol := "http"
		if (*coreinfo)["ssl"] == "true" {
			protocol = "https"
		}

		port, _ := (*coreinfo)["port"].(float64)
		fmt.Printf("  %-25s %s\n", "OS Version:", (*hostinfo)["operating_system"])
		fmt.Printf("  %-25s %s\n", "Muthur Command Core:", (*coreinfo)["version"])
		fmt.Println()
		fmt.Printf("  %-25s %s://%s.local:%d\n", "Muthur Command URL:", protocol, (*hostinfo)["hostname"], int(port))
		fmt.Printf("  %-25s http://%s.local:%d\n", "Observer URL:", (*hostinfo)["hostname"], 4357)
		fmt.Println("")
		fmt.Println("System is ready! Use browser or app to configure.")
	},
}

func init() {
	rootCmd.AddCommand(bannerCmd)
	bannerCmd.Flags().Bool("no-wait", false, "Don't wait until Supervisor is started")
}
