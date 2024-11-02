package database

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./employees.db")
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	createJobsTable := `
	CREATE TABLE IF NOT EXISTS jobs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		employee_id INTEGER NOT NULL,
		department TEXT NOT NULL,
		job_title TEXT NOT NULL,
		FOREIGN KEY (employee_id) REFERENCES employees(id)
	);`

	if _, err := DB.Exec(createJobsTable); err != nil {
		log.Fatalf("Error creating jobs table: %v", err)
	}
}
