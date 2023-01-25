package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Jewels2001/seekers_guild/api/db"
	"github.com/Jewels2001/seekers_guild/api/util"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[REQUEST] \"/register\" POST\t{%s}\n", r.RemoteAddr)

	// Extract incoming data
	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Println("[REQUEST] bad body")
		util.RespondWithError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	name, ok := body["name"]
	if !ok {
		log.Println("[REQUEST] bad body")
		util.RespondWithError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	email, ok := body["email"]
	if !ok {
		log.Println("[REQUEST] bad body")
		util.RespondWithError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	passwdHash, ok := body["passwdHash"]
	if !ok {
		log.Println("[REQUEST] bad body")
		util.RespondWithError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	// Add user to DB
	id, err := db.AddUser(name, email, passwdHash)
	if err != nil {
		log.Println("[ERROR]", err.Error())
		util.RespondWithError(w, http.StatusInternalServerError, "error adding user")
		return
	}

	// Check to see if user was inserted
	if id == -1 {
		log.Println("[REQUEST] unable to create duplicate user")
		util.RespondWithError(w, http.StatusConflict, "user already exists")
		return
	}

	// Return user
	full_user, err := db.GetUser(id)
	if err != nil {
		log.Println("[ERROR]", err.Error())
		util.RespondWithError(w, http.StatusInternalServerError, "error retreiving user")
		return
	}

	log.Println("[REQUEST] added user:", id)
	util.RespondWithJSON(w, http.StatusOK, full_user)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[REQUEST] \"/login\" POST\t{%s}\n", r.RemoteAddr)

}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[REQUEST] \"/logout\" POST\t{%s}\n", r.RemoteAddr)

}
