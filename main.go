package main

import (
	"log"
	"os"
	"search-by-ip/app"
)

func main() {
	app := app.Generate()
	error := app.Run(os.Args)

	if error != nil {
		log.Fatal(error)
	}
}