package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Load_DB(dsn string) {
	var err error
	DB, err = sql.Open("sqlite", dsn)
	if err != nil {
		log.Fatalf("[Error] failed to open database err: %v", err)
	}

	schema := `
	CREATE TABLE IF NOT EXISTS tasks(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			status TEXT CHECK(status in ('pending','running','done')),
			result TEXT,
			delay_time INTEGER,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)
	`

	_, err = DB.Exec(schema)
	if err != nil {
		log.Fatalf("[Error]failed to init database err:%v", err)
	}

	log.Printf("[Init] database init at %s", dsn)
}
