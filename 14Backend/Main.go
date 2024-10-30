// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/go-sql-driver/mysql"
// )

// func initDB() *sql.DB {
// 	dsn := "root:Viveksequel24@tcp(127.0.0.1:3306)/viv1"
// 	// Here dsn is the string that acts as a data source name which basically gives necessary information that helps to coonect with the database.
// 	db, err := sql.Open("mysql", dsn)

// 	// sql.Open() function takes 2 args , first one is the name of the driver in this case it is mysql whereas 2nd parameter is a dsn string.

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Successfully connected to datbase")

// 	return db
// }

// func CreateUser(db *sql.DB, name string, age int) {
// 	query := "INSERT INTO users(name,age)values(?,?)"

// 	result, err := db.Exec(query, name, age)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	id, _ := result.LastInsertId()
// 	fmt.Printf("New user created with the id %d\n", id)

// }

// func UpdateUser(db *sql.DB, UserID int, NewAge int) {
// 	query := "UPDATE users SET age= ? WHERE id = ?"

// 	_, err := db.Exec(query, UserID, NewAge)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("User Updated with age :- %d having id = %d", NewAge, UserID)
// }

// func main() {
// 	db := initDB()
// 	defer db.Close()

// 	// CreateUser(db, "Penny", 22)
// 	UpdateUser(db, 2, 34)
// }

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

	return db

}

func CreateUser(db *sql.DB, name string, age int) {

	query := "INSERT INTO users(name,age)values(?,?)"

	result, err := db.Exec(query, name, age)

	if err != nil {
		log.Fatal(err)
	}

	id, _ := result.LastInsertId()
	fmt.Printf("New user added with id %d", id)

}

func main() {

	fmt.Println("Hello, I am learning to integrate Database to my backend")

	db := initDB()

	CreateUser(db, "mahesh", 34)

}
