package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func initDB() *sql.DB {
	dsn := "root:Viveksequel24@tcp(127.0.0.1:3306)/viv1"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MYSQL Database Successfully")

	return db
}

func CreateUser(db *sql.DB, name string, age int) {
	query := "INSERT INTO users(name,age)VALUES(?,?)"

	_, err := db.Exec(query, name, age)

	if err != nil {
		log.Fatal(err)
	}
	// result.LastInsertId()
	// fmt.Printf("New user inserted with id :- %d\n", id)
}

func main() {
	db := initDB()
	defer db.Close()
	CreateUser(db, "alice", 30)
}
