package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
  pool, err := sql.Open("sqlite3", "api.db")

  if err != nil {
    panic("Could not connect to the database.")
  }

  DB = pool

  DB.SetMaxOpenConns(10)
  DB.SetMaxIdleConns(5)

  createTables()
}

func createTables() {
  createUsersTableQuery := `
  CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL
  )
  `

  _, err := DB.Exec(createUsersTableQuery)

  if err != nil {
    panic("Could not create users table.")
  }

  createEventsTableQuery := `
  CREATE TABLE IF NOT EXISTS events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    location TEXT NOT NULL,
    dateTime DATETIME NOT NULL,
    user_id INTEGER,
    FOREIGN KEY(user_id) REFERENCES users(id)
  )
  `

  _, err = DB.Exec(createEventsTableQuery)

  if err != nil {
    panic("Could not create events table.");
  }
}
