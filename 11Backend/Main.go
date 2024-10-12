package main

import (
	"fmt"
	"net/http"
	"os"
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

	// var home string = "/home"

	var form string = "/form"

	fs := http.FileServer(http.Dir("./Static"))

	http.Handle(form, http.StripPrefix(form, fs))

	http.HandleFunc(root, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the root path of my server !")
	})

	http.HandleFunc(user, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "user accessing this page is :- %s %s", user1.firstname, user1.lastname)
	})

	// http.HandleFunc(home, func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "welcome home")
	// })

	/* You’ll add a new route, /submit, to handle the form data when it’s submitted.
	The handler will read the form values from the request and respond accordingly. */

	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		// check if the form is submitted with post request
		if r.Method != http.MethodPost {
			http.Error(w, "An error has occurred :-", http.StatusMethodNotAllowed)
			return
		}

		//parse the form data

		if err := r.ParseForm(); err != nil {
			http.Error(w, "An error occurred :-", http.StatusInternalServerError)
			return
		}

		// Retrieve form values

		Username := r.FormValue("username")
		Password := r.FormValue("password")
		DOB := r.FormValue("date")

		file, err := os.OpenFile("./Static/Users.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)

		if err != nil {
			fmt.Println("error appending new user to the users file", err)
		}

		file.WriteString(Username)
		file.WriteString(Password)
		file.WriteString(DOB)

		fmt.Fprintf(w, "Recieved form data successfully with the following inputs :- \n")
		fmt.Fprintf(w, "Username:- %s\n", Username)
		fmt.Fprintf(w, "Password:- %s\n", Password)
		fmt.Fprintf(w, "Date of Birth:- %s\n", DOB)

	})

	var port string = ":8080"
	fmt.Println("Server is up and running at :- ", port)
	http.ListenAndServe(port, nil)

}
