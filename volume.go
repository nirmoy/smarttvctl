package main

import (
	"flag"
	"fmt"
	"net/url"
	"strings"

	"github.com/huin/goupnp"
	"github.com/huin/goupnp/dcps/av1"
)

func main() {
	vol := flag.Int("vol", 42, "Volume from 0-100")
	flag.Parse()
	streamingService := map[int]*url.URL{}
	//allService, _ := goupnp.DiscoverDevices("ssdp:all")
	allService, _ := goupnp.DiscoverDevices(av1.URN_RenderingControl_1)
	fmt.Println("Discovered:")
	for _, service := range allService {
		fmt.Println(service.Root.Device.Manufacturer)
		if strings.Contains(service.Root.Device.Manufacturer, "Samsung") {
			streamingService[1] = service.Location
			fmt.Println(service.Root.Device.Manufacturer)
			break
		}
	}
	avTransports, err := av1.NewRenderingControl1ClientsByURL(streamingService[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, av := range avTransports {
		fmt.Println(av.SetVolume(0, "Master", uint16(*vol)))
	}

}
