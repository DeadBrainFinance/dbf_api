package main

import (
	"log"

	"dbf_api/api"
)

func main() {
	app := api.NewApp()
	port := "4000"

	if err := app.Run(port); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
