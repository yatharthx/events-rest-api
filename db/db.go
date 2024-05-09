package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
  pool, err := sql.Open("sqlite3", "api.db")

  if err != nil {
    panic("Could not connect to the database.")
  }

  DB = pool
  // defer DB.Close()

  DB.SetMaxOpenConns(10)
  DB.SetMaxIdleConns(5)

  createTables()
}

func createTables() {
  createEventsTableQuery := `
  CREATE TABLE IF NOT EXISTS events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    location TEXT NOT NULL,
    dateTime DATETIME NOT NULL,
    user_id INTEGER
  )
  `

  _, err := DB.Exec(createEventsTableQuery)

  if err != nil {
    fmt.Printf("%v", err);
  }
}
