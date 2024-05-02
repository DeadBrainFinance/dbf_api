package main

import (
	"log"

	"dbf_api/server"
)

func main() {
	app := server.NewApp()
	port := "4000"

	if err := app.Run(port); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
