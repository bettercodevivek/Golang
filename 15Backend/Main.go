package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Book struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

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

func main() {

	db := initDB()

	defer db.Close()

}
