package main

import (
	"fmt"
	"go-new-http-web/routes"
)

func main() {
	app := routes.Initialize()
	app.LoadRoutes()
	app.Run()

	fmt.Println("Server is running on port 8080...")
}
