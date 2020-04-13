package mssql

import (
	"database/sql"
	"log"
)

// NewDBConnection create new connection to mssql using the specified connection string
func NewDBConnection(connString string) (*sql.DB, error) {
	db, err := sql.Open("mssql", connString)

	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
		return nil, err
	}

	return db, nil
}
