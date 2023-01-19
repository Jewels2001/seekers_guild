package main

import (
    "log"
    "github.com/Jewels2001/seekers_guild/api/db"
)

func main() {
    log.Println("TEST")
    if err := db.InitDB(); err != nil {
        log.Fatal(err)
    }
    defer db.ShutdownDB()
}
