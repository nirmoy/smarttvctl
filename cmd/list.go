package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/huin/goupnp"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var (
	allDevices bool
)
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "list smart TVs in the network",
	//Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		deviceList := map[string][]string{}
		allService, _ := goupnp.DiscoverDevices("ssdp:all")
		for _, service := range allService {
			if service.Root == nil {
				continue
			}
			if !allDevices {
				if strings.Contains(service.Location.String(), "dmr") {
					deviceList[service.Location.String()] =
						[]string{service.Root.Device.Manufacturer, service.Location.String()}
				}
			} else {
				deviceList[service.Location.String()] =
					[]string{service.Root.Device.FriendlyName, service.Location.String()}
			}
		}

		fmt.Println("Discovered:")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "URL"})

		for _, val := range deviceList {
			table.Append(val)
		}

		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&allDevices, "all", "a", false, "List all devices")
}
