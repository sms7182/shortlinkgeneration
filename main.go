package main

import (
	"log"
	"net/http"
	"os"
	"shortlinkapi/app"
	"shortlinkapi/database"
)

func main() {
	app := app.New()
	app.DB = &database.DB{}

	http.HandleFunc("/", app.Router.ServeHTTP)

	err := http.ListenAndServe(":9000", nil)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
