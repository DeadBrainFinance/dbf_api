package main

import (
	"context"
	"fmt"
	
	"dbf_api/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}

