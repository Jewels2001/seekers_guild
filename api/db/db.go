package db

import (
	"database/sql"
	"errors"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const DB_PATH = "./db/seekers_guild.db"

var db *sql.DB
var err error

func InitDB() error {
	if _, err := os.Stat(DB_PATH); errors.Is(err, os.ErrNotExist) {
		log.Println("DB does not exist, creating...")
	}

	db, err = sql.Open("sqlite3", DB_PATH)
	if err != nil {
		return err
	}

    _, err = db.Exec(create_users_table)
    if err != nil {
        return err
    }

	return nil
}

func ShutdownDB() {
    log.Println("closing db...")
    db.Close()
}
