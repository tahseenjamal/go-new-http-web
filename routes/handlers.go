package routes

import (
	"fmt"
	"net/http"

	"go-new-http-web/controllers"
)

type App struct {
	Router *http.ServeMux
}

func Initialize() *App {
	app := &App{
		Router: http.NewServeMux(),
	}

	return app
}

func (app *App) Run(port ...int) {
	var listenerPort string

	if len(port) == 0 {
		listenerPort = "8081"
	}

	v1 := http.NewServeMux()

	// Versioning implementation using http.StripPrefix
	v1.Handle("/v1/", http.StripPrefix("/v1", app.Router))

	http.ListenAndServe(fmt.Sprintf(":%s", listenerPort), v1)
}

func (app *App) LoadRoutes() {
	user := &controllers.User{}

	app.Router.HandleFunc("GET /user", user.GetAll)
	app.Router.HandleFunc("GET /user/{id}", user.Get)
	app.Router.HandleFunc("POST /user", user.Post)
	app.Router.HandleFunc("PUT /user/{id}", user.Put)
	app.Router.HandleFunc("DELETE /user/{id}", user.Delete)
}
