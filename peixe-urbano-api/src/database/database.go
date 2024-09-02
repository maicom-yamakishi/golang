package database

import (
	"database/sql"
	"peixe-urbano/src/config"

	_ "github.com/go-sql-driver/mysql" //Driver
)

// Connect open connection with database and return
func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.ConnectionString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
