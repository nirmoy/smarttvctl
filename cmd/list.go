package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/huin/goupnp"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
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
			if strings.Contains(service.Location.String(), "dmr") {
				deviceList[service.Root.Device.Manufacturer] =
					[]string{service.Root.Device.Manufacturer, service.Location.String()}
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
}
