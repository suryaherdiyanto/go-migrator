package gomigrator

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func NewConnection(dataSourceName string) error {
	DB, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		log.Printf("Could not make connection to the database: %v", err)
		return err
	}

	if DB.Ping() != nil {
		err = errors.New("could not ping to database")
		log.Print(err)
		return err
	}

	return nil
}
