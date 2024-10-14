package main

import (
	"database/sql"
	"fmt"

	_ "github.com/ncruces/go-sqlite3/driver"

	_ "github.com/ncruces/go-sqlite3/embed"
)

const setupDB = `
	CREATE TABLE IF NOT EXISTS album (
		id         INTEGER PRIMARY KEY AUTOINCREMENT,
		title      VARCHAR(128) NOT NULL,
		artist     VARCHAR(255) NOT NULL,
		price      DECIMAL(5,2) NOT NULL
	);
`

func setupDb(dbName string) *sql.DB {
	var err error

	db, err := sql.Open("sqlite3", fmt.Sprintf("file:%v", dbName))
	if err != nil {
		errorLog.Fatal(err)
	}
	_, err = db.Exec(setupDB)
	if err != nil {
		errorLog.Error(err)
	} else {
		infoLog.Info("Database initialised successfully")
	}
	return db
}

func closeDb(db *sql.DB) {
	db.Close()
}
