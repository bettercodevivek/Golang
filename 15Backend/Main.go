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

func initDB() {
	var err error
	// Corrected DSN
	dsn := "root:Viveksequel24@tcp(127.0.0.1:3306)/viv1"

	// Assign to the global variable
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	// Ensure the connection works
	if err := db.Ping(); err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	fmt.Println("Connection established successfully with DB")
}

// Get All Books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, title, author, price FROM books")
	if err != nil {
		http.Error(w, "Failed to fetch books", http.StatusInternalServerError)
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

	// Set response header and send the JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func AddBook(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid http Method", http.StatusBadRequest)
		return
	}

	var newBook Book

	err := json.NewDecoder(r.Body).Decode(&newBook)

	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO books(title,author,price)values(?,?,?)"

	result, err := db.Exec(query, newBook.Title, newBook.Author, newBook.Price)

	if err != nil {
		http.Error(w, "Failed to insert book", http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	newBook.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newBook)

}

func main() {
	initDB()
	defer db.Close()

	// Register the handler
	http.HandleFunc("/books", GetBooks)

	http.HandleFunc("/addbook", AddBook)

	// Start the server
	fmt.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
