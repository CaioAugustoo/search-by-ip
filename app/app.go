package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func searchIps(context *cli.Context) {
	host := context.String("host")

	ips, error := net.LookupIP(host)

	if error != nil {
		log.Fatal(error)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServers(context *cli.Context) {
	host := context.String("host")

	servers, error := net.LookupNS(host)

	if error != nil {
		log.Fatal(error)
	}

	for _, server := range servers {
		fmt.Println(server)
	}
}

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Search IP's"
	app.Usage = "Search IP's and server names from web"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "caioaugusto.dev",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Search IP's from web",
			Flags: flags,
			Action: searchIps,
		},
		{
			Name: "server",
			Usage: "Search name servers from web",
			Flags: flags,
			Action: searchServers,
		},
	}

	return app
}