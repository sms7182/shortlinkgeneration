package app

import (
	"shortlinkapi/database"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *database.DB
}

func New() *App {
	a := &App{
		Router: mux.NewRouter(),
	}

	a.initRoutes()
	return a

}

func (a *App) initRoutes() {
	a.Router.HandleFunc("/", a.IndexHandler()).Methods("GET")
}
