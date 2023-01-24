package routes

import (
	"log"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[REQUEST] \"/register\" POST\t{%s}\n", r.RemoteAddr)

          
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[REQUEST] \"/login\" POST\t{%s}\n", r.RemoteAddr)

}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[REQUEST] \"/logout\" POST\t{%s}\n", r.RemoteAddr)

}
