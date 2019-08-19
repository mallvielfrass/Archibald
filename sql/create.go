package main

import (
	"database/sql"
	//	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	//	"os"
)

func main() {

	db, err := sql.Open("sqlite3", "./user.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
//id username first_name last_name
	sqlStmt := `
	CREATE TABLE data(
  		id    		INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE, 
  		username  	TEXT,
  		first_name	TEXT,
  		last_name	TEXT
	)	
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}
