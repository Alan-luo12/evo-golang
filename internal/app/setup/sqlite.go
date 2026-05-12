package setup

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

// 初始化SQLite数据库
func InitSqliteDB(dsn string) *sql.DB {
	var db *sql.DB
	//打开SQLite数据库
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		log.Fatal("[Error] Failed to open database")
	}

	//创建表
	schema := `
	CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER PRIMARY KEY ,
			name TEXT NOT NULL,
			result TEXT,
			status TEXT CHECK(status IN ('running','done','failed')),
			delay_time INTEGER,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)
	`

	//执行创建表的SQL语句
	_, err = db.Exec(schema)
	if err != nil {
		log.Fatalf("[Error] Failed to init database at %s", dsn)
	}

	log.Printf("[Init] success to init at %s", dsn)
	return db
}
