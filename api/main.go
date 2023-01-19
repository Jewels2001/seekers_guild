package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Jewels2001/seekers_guild/api/db"

	"github.com/gorilla/mux"
)

const (
    ADDR = "localhost"
    PORT = "8080"
)

// Hello world landing page for root endpoint
//
func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "SEEKERS_GUILD API is ONLINE: %s", time.Now().String())
}

func main() {
    // Input flags
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

    // Setup DB
	if err := db.InitDB(); err != nil {
		log.Fatal(err)
	}
	defer db.ShutdownDB()

    // Setup Router
	r := mux.NewRouter()
	r.HandleFunc("/", RootHandler)

    // Set timeouts on connections
	srv := &http.Server{
        Addr:         ADDR+":"+PORT,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

    // Serve API in goroutine to avoid blocking
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
    log.Printf("server running at: http://%s:%s", ADDR, PORT)

    // Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
    log.Println("shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	srv.Shutdown(ctx)
	log.Println("shutting down")
}