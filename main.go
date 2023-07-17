package main

import (
	"flag"

	"github.com/codescalersinternships/envserver-Diaa/server"
)

func main(){
	portFlag := flag.Int("p",8080,"port that will be used to run the app")

	flag.Parse()

	app:=server.App{}
	app.SetPort(*portFlag)
	app.Run()
}