package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"

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

	// Extract incoming data
	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
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

	// Get user from db
	user, err := db.GetUserByEmail(email)
	if err != nil {
		log.Println("[ERROR]", err.Error())
		util.RespondWithError(w, http.StatusInternalServerError, "unable to retreive user")
		return
	}

	// Check password hash against db
	if user.ValidatePasswordHash(passwdHash) {
		log.Printf("[REQUEST] user %d successfully logged in\n", user.Id)

		// Generate token
		tokenParams := db.Token{
			Uid: user.Id,
			Aid: uuid.New().String(),
		}
		claims := map[string]string{
			"uid": strconv.Itoa(tokenParams.Uid),
			"aid": tokenParams.Aid,
		}

		token, err := util.GenerateToken(claims, "HS256")
		if err != nil {
			log.Println("[ERROR] error generating token:", err.Error())
			util.RespondWithError(w, http.StatusInternalServerError, "error generating token")
			return
		}

		// Add token to db
		if _, err = db.AddToken(tokenParams); err != nil {
			log.Println("[ERROR] error writing token to database")
			util.RespondWithError(w, http.StatusInternalServerError, "error generating token")
			return
		}

		util.RespondWithJSON(w, http.StatusOK, map[string]string{"token": token})
	} else {
		log.Printf("[REQUEST] invalid login attempt for user %d\n", user.Id)
		util.RespondWithError(w, http.StatusUnauthorized, "incorrect credentials")
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[REQUEST] \"/logout\" POST\t{%s}\n", r.RemoteAddr)

	// Extract incoming data
	var body map[string]int
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Println("[REQUEST] bad body")
		util.RespondWithError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	id, ok := body["id"]
	if !ok {
        log.Println("[REQUEST] bad body: no id")
		util.RespondWithError(w, http.StatusBadRequest, "invalid request body")
		return
	}

    // Remove all tokens associated with uid
    if err := db.RemoveToken(id); err != nil {
        log.Println("[ERROR] error invalidating tokens:", err.Error())
		util.RespondWithError(w, http.StatusInternalServerError, "error invalidating tokens")
		return
    }
    
    log.Printf("[REQUEST] user %d successfully logged out\n", id)
    util.RespondWithJSON(w, http.StatusOK, map[string]string{"response":"successfully logged out"})
}
