package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Create() *cli.App {
	app := cli.NewApp()
	app.Name = "Command Line Aplication"
	app.Usage = "Search IPs and Server Names on Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "amazon.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "IP's search on Internet",
			Flags:  flags,
			Action: searchIPs,
		},

		{
			Name:   "servers",
			Usage:  "Servers' search on Internet",
			Flags:  flags,
			Action: searchServers,
		},
	}

	return app
}

func searchIPs(c *cli.Context) {
	host := c.String("host")

	ips, error := net.LookupIP(host)
	for error != nil {
		log.Fatal(error)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServers(c *cli.Context) {
	host := c.String("host")

	servers, error := net.LookupNS(host)
	if error != nil {
		log.Fatal(error)
	}
	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
