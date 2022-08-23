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

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Search IP's"
	app.Usage = "Search IP's and server names from web"

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Search IP's from web",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "host",
					Value: "caioaugusto.dev",
				},
			},

			Action: searchIps,
		},
	}

	return app
}