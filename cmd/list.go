package cmd

import (
	"fmt"
	"strings"

	"github.com/huin/goupnp"
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
		for key, val := range deviceList {
			fmt.Println(key, val)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
