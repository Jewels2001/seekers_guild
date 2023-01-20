package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Jewels2001/seekers_guild/api/db"
	"github.com/Jewels2001/seekers_guild/api/util"
	"github.com/gorilla/mux"
)

func PrestigeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Printf("[REQUEST] \"/users/%s/updatePrestige\" PATCH\t{%s}\n", vars["id"], r.RemoteAddr)

	// Parse params
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Println("[REQUEST] invalid request")
		util.RespondWithError(w, http.StatusBadRequest, "invalid reqest")
		return
	}

	// Extract body
	body := make(map[string]float64, 0)
	if err = json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Println("[REQUEST] invalid request")
		util.RespondWithError(w, http.StatusBadRequest, "invalid reqest")
		return
	}
	changeAmount, ok := body["prestige_change"]
	if !ok {
		log.Println("[REQUEST] invalid request")
		util.RespondWithError(w, http.StatusBadRequest, "invalid reqest")
		return
	}
  //   changeAmount, err := strconv.ParseFloat(data, 64)
  //   if err != nil {
		// log.Println("[REQUEST] invalid request")
		// util.RespondWithError(w, http.StatusBadRequest, "invalid reqest")
		// return
  //   }

	// Execute update query
	newPrestige, err := db.UpdatePrestige(int(id), changeAmount)
	if err != nil {
		log.Println("[ERROR]", err.Error())
		util.RespondWithError(w, http.StatusInternalServerError, "unable to update prestige")
		return
	}

	log.Printf("[REQUEST] prestige updated for user: %d (new value: %0.2f)\n", id, newPrestige)
	util.RespondWithJSON(w, http.StatusOK, map[string]string{
		"message":  "prestige successfully updated",
		"prestige": strconv.FormatFloat(newPrestige, 'f', 2, 32),
	})
}
