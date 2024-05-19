package main

import "go-new-http-web/routes"

func main() {
	app := routes.Initialize()
	app.LoadRoutes()
	app.Run()
}
