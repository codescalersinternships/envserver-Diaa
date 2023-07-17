package main

import (
	"flag"
	"github.com/codescalersinternships/envserver-Diaa/server"
	"log"
)

func main() {
	var portFlag int
	flag.IntVar(&portFlag, "p", 8080, "port that will be used to run the app")

	flag.Parse()

	app := server.NewApp()
	err := app.SetPort(portFlag)
	if err != nil {
		log.Fatal(err)
	}

	err = app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
