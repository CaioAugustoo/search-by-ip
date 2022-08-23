package main

import (
	"log"
	"os"
	"search-by-ip/app"
)

func main() {
	createdCli := app.Generate()
	error := createdCli.Run(os.Args)

	if error != nil {
		log.Fatal(error)
	}
}