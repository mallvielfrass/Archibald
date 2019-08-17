package main

import (
	"database/sql"
	//	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	//	"os"
)

func main() {

	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	CREATE TABLE data(
  		id    INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE, 
  		username  TEXT,
  		password  TEXT,
  		token TEXT
	)	
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}
