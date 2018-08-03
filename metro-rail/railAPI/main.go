package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/uzumakitensho/chapter04/metro-rail/dbutils"
	"log"
)

func main() {
	db, err := sql.Open("sqlite3", "./railapi.db")
	if err != nil {
		log.Println("Driver creation failed!")
	}
	dbutils.Initialize(db)
}
