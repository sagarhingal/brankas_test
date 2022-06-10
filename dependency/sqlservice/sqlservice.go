package sqlservice

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

// Initdatabase : This function initialises the database
// and creates a shared object for the other services to
// use.
func Initdatabase(filename string, requestID int64) {
	var err error
	os.Remove(filename)
	Db, err = sql.Open("sqlite3", filename)
	if err != nil {
		log.Fatal("[", requestID, "] | Initdatabase() - Unable to initialise SQLite3. | ", err)
	} else {
		if err := createtable(); err != nil {
			log.Fatalln("[", requestID, "] | Initdatabase() - Unable to create table. | ", err)
		} else {
			log.Println("[", requestID, "] | Initdatabase() - SQLite initialised successfully.")
		}

	}

}

// createtable : This functions creates a new table upon invocation.
func createtable() error {
	query := `
	CREATE TABLE IF NOT EXISTS imgmeta(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		originalname TEXT NOT NULL,
		newname TEXT NOT NULL UNIQUE,
		filesize INTEGER NOT NULL,
		contenttype TEXT NOT NULL,
		agent TEXT NOT NULL,
		clientip TEXT NOT NULL,
		createdat DATETIME DEFAULT CURRENT_TIMESTAMP,
		updatedat DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := Db.Exec(query)
	return err
}
