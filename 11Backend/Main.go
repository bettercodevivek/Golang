package main

import (
	"fmt"
	"net/http"
)

type user struct {
	firstname string
	lastname  string
	age       int
}

func main() {

	user1 := user{"vivek", "singh", 22}

	var root string = "/"

	var user string = "/user"

	http.HandleFunc(root, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the root path of my server !")
	})

	http.HandleFunc(user, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "user accessing this page is :- "+user1.firstname+user1.lastname)
	})

	var port string = ":8080"
	fmt.Println("Server is up and running at :- ", port)
	http.ListenAndServe(port, nil)

}
