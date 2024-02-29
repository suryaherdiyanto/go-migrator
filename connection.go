package gomigrator

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func New() {
	DB, err := sql.Open("mysql", "root:root@/testdb")

	if err != nil {
		log.Fatalf("Could not make connection to the database: %v", err)
	}

	if DB.Ping() != nil {
		log.Fatalf("Error ping to database: %v", err)
	}
}
