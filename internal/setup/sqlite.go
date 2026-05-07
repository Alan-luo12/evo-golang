package setup

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func InitDB(dsn string) *sql.DB {
	var db *sql.DB
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		log.Fatal("[Error] Failed to open database")
	}

	schema := `
	CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER PRIMARY KEY ,
			name TEXT NOT NULL,
			result TEXT,
			status TEXT CHECK(status IN ('pending','running','done')),
			delay_time INTEGER,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)
	`

	_, err = db.Exec(schema)
	if err != nil {
		log.Fatalf("[Error] Failed to init database at %s", dsn)
	}

	log.Printf("[Init] success to init at %s", dsn)
	return db
}
