package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Jewels2001/seekers_guild/api/db"
	"github.com/Jewels2001/seekers_guild/api/util"
	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[REQUEST] \"/users\" GET\t{%s}\n", r.RemoteAddr)

	// Get users from DB
	users, err := db.GetUsers()
	if err != nil {
		log.Println("[ERROR]", err.Error())
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Printf("[REQUEST] \"/users/%s\"\t{%s}\n", vars["id"], r.RemoteAddr)

	// Parse params
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Println("[REQUEST] invalid request")
		util.RespondWithError(w, http.StatusBadRequest, "invalid reqest")
		return
	}

	// Get user from DB
	user, err := db.GetUser(int(id))
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			log.Println("[REQUEST] user not found")
			util.RespondWithError(w, http.StatusNotFound, fmt.Sprintf("id: %d not found", id))
		} else {
			log.Println("[ERROR]", err.Error())
			util.RespondWithError(w, http.StatusInternalServerError, "internal error")
		}
		return
	}

	util.RespondWithJSON(w, http.StatusOK, user)
}

func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[REQUEST] \"/users\" POST\t{%s}\n", r.RemoteAddr)

	// Extract incoming data
	var u db.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		log.Println("[REQUEST] bad body")
		util.RespondWithError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	// Add user to DB
	res, err := db.AddUser(u)
	if err != nil {
		log.Println("[ERROR]", err.Error())
		util.RespondWithError(w, http.StatusInternalServerError, "error adding user")
		return
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Println("[ERROR]", err.Error())
		util.RespondWithError(w, http.StatusInternalServerError, "error retreiving user")
		return
	}
	full_user, err := db.GetUser(int(id))
	if err != nil {
		log.Println("[ERROR]", err.Error())
		util.RespondWithError(w, http.StatusInternalServerError, "error retreiving user")
		return
	}

	log.Println("[REQUEST] added user:", id)
	util.RespondWithJSON(w, http.StatusOK, full_user)
}

func RemoveUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Printf("[REQUEST] \"/users/%s\" DELETE\t{%s}\n", vars["id"], r.RemoteAddr)

	// Parse params
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Println("[REQUEST] invalid request")
		util.RespondWithError(w, http.StatusBadRequest, "invalid reqest")
		return
	}

	// Make transaction in db
	if err = db.RemoveUser(int(id)); err != nil {
		log.Println("[ERROR] error removing user")
		util.RespondWithError(w, http.StatusInternalServerError, "error removing user")
		return
	}

	log.Println("[REQUEST] removed user:", id)
	util.RespondWithJSON(w, http.StatusOK, "user successfully removed")
}
