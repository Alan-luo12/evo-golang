package setup

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func InitMysqlDB(dsn string, maxconns int, maxidleconn int, connmaxlifetime time.Duration) *sql.DB {
	var db *sql.DB
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("[Error] Failed to open database at %s", dsn)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("[Error] Failed to ping database at %s", dsn)
	}

	db.SetMaxOpenConns(maxconns)
	db.SetMaxIdleConns(maxidleconn)
	db.SetConnMaxLifetime(connmaxlifetime)

	schema := `
		CREATE TABLE IF NOT EXISTS tasks (
			id BIGINT PRIMARY KEY ,
			name TEXT NOT NULL,
			result TEXT,
			status TEXT CHECK(status IN ('running','done','failed')),
			delay_time INTEGER,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
		`

	_, err = db.Exec(schema)
	if err != nil {
		log.Fatalf("[Error] Failed to create table tasks")
		return nil
	}

	log.Printf("[Info] Database connection established")
	return db
}
