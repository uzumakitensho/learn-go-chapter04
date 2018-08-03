package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Book struct {
	id     int
	name   string
	author string
}

func main() {
	db, err := sql.Open("sqlite3", "./books.db")
	log.Println(db)
	if err != nil {
		log.Println(err)
	}

	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS books (id INTEGER PRIMARY KEY, isbn INTEGER, author VARCHAR(64), name VARCHAR(64) NULL)")
	if err != nil {
		log.Println("Error in creating table")
	} else {
		log.Println("Successfully created table books!")
	}
	statement.Exec()

	statement, _ = db.Prepare("INSERT INTO books (name, author, isbn) VALUES (?,?,?)")
	statement.Exec("A Tale of Two Towns", "Bangkeee Ayaaamm", 12345678)
	log.Println("Inserted the book into database!")

	rows, _ := db.Query("SELECT id, name, author FROM books")
	var tempBook Book
	for rows.Next() {
		rows.Scan(&tempBook.id, &tempBook.name, &tempBook.author)
		log.Printf("ID:%d, Book:%s, Author:%s\n", tempBook.id, tempBook.name, tempBook.author)
	}

	statement, _ = db.Prepare("UPDATE books set name=? where id=?")
	statement.Exec("Judul hmmmm..", 1)
	log.Println("Successfully updated the book in database!")

	statement, _ = db.Prepare("DELETE from books where id=?")
	statement.Exec(1)
	log.Println("Successfully deleted the book in database!")
}
