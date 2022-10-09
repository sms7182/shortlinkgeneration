package app

import (
	"net/http"
	"shortlinkapi/database"
	"strings"

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
func CaselessMatcher(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.ToLower(r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func (a *App) initRoutes() {
	CaselessMatcher(a.Router)
	a.Router.HandleFunc("/", a.IndexHandler()).Methods("GET")
	a.Router.HandleFunc("/sendurl", a.SendUrlHanlder()).Methods("POST")
	a.Router.HandleFunc("/{urlShortCode}", a.RedirectIfURLExist).Methods("GET")
	a.Router.HandleFunc("/redirect", a.TestRedirect()).Methods("POST")
}
