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

	http.ListenAndServe(fmt.Sprintf(":%s", listenerPort), app.Router)
}

func (app *App) LoadRoutes() {
	user := &controllers.User{}

	app.Router.HandleFunc("GET /user", user.GetAll)
	app.Router.HandleFunc("GET /user/{id}", user.Get)
	app.Router.HandleFunc("POST /user", user.Post)
	app.Router.HandleFunc("PUT /user/{id}", user.Put)
	app.Router.HandleFunc("DELETE /user/{id}", user.Delete)
}
