package gomigrator

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func NewConnection(driverName, dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open(driverName, dataSourceName)

	if err != nil {
		log.Printf("Could not make connection to the database: %v", err)
		return db, err
	}

	if db.Ping() != nil {
		err = errors.New("could not ping to database")
		log.Print(err)
		return db, err
	}

	return db, nil
}
