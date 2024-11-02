package main

import (
	"15Backend_MyApp/api/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {

	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "WELCOME TO THE SERVER")
	})

	router.HandleFunc("/api/users", handlers.UsersHandler)

	router.HandleFunc("/api/users/", handlers.UserHandler)

	log.Println("Server running on port:8080")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}

}
