package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func UserRegisterHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("User registriation hit")
	var newUser User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(400)
		log.Error("User creation failed with ", err)
		return
	}
	prunus.createUser(newUser.Username, newUser.Password)
	fmt.Fprintf(w, "{\"error\":None}")


	log.Info("Created user", newUser.Username)
}

func GetUsersHandler(w http.ResponseWriter, _ *http.Request) {
	log.Info("Get user info hit")
	data, err := json.Marshal(prunus.Users)
	if err != nil {
		log.Error("JSON Marshal error: ", err)
	}

	fmt.Fprint(w, string(data))
}
