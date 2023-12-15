package main
import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)
var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "car_garage.db")
	if err != nil {
		log.Fatal(err)
	}
	createTable()
}

func createTable() {
	query := `
		CREATE TABLE IF NOT EXISTS cars (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			brand TEXT,
			model TEXT,
			status TEXT
		);
	`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}