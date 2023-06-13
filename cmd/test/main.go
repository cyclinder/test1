package main

import (
	"flag"
	"github.com/containernetworking/plugins/pkg/ns"
	"github.com/vishvananda/netlink"
	"log"
)

var nsPath = flag.String("path", "", "ns path")

func main() {
	flag.Parse()

	if *nsPath == "" {
		log.Fatalln("path can't be empty")
	}

	netns, err := ns.GetNS(*nsPath)
	if err != nil {
		log.Fatalln(err.Error())
	}

	var routes []netlink.Route
	err = netns.Do(func(_ ns.NetNS) error {
		routes, err = netlink.RouteList(nil, netlink.FAMILY_ALL)
		return err
	})

	if err != nil {
		log.Fatalln(err)
	}

	for _, route := range routes {
		log.Println(route.String())
	}

}
