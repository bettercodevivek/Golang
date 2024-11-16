package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Book struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

var db *sql.DB

func initDB() *sql.DB {
	dsn := "root:Viveksequel24@tcp(127.0.0.3306)/viv1"

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal(err)
	}

	db.Ping()

	fmt.Println("Connection established successfully with DB")

	return db

}

// Get All Books

func GetBooks(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id,title,author,price from books")

	if err != nil {
		http.Error(w, "Failed to Fetch Books", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var books []Book

	for rows.Next() {
		var book Book

		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price); err != nil {
			http.Error(w, "Error scanning books", http.StatusInternalServerError)
			return
		}
		books = append(books, book)
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func main() {

	db := initDB()

	defer db.Close()

}
