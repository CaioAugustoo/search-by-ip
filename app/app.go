package app

import "github.com/urfave/cli"

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Search IP's"
	app.Usage = "Search IP's and server names from web"

	return app
}