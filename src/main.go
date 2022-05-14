package main

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var (
	prunus DB
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/register", UserRegisterHandler)
	r.HandleFunc("/users", GetUsersHandler)
	log.Info("Started Webserver.")
	http.ListenAndServe(":8080", r)
}
