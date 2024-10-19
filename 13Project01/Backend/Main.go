package main

import (
	"fmt"
	"net/http"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"Password"`
	DOB      string `json:"DOB"`
}

var userFile = "users.json"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to vivek's server,please direct to /form route to enter your details and register")
	})

	http.HandleFunc("/form/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../Frontend/index.html")
	})

	fs := http.FileServer(http.Dir("../Frontend/Static"))

	http.Handle("/Static/", http.StripPrefix("/Static/", fs))

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Error occured in method :- ", http.StatusMethodNotAllowed)
			return
		}

		if err := r.ParseForm(); err != nil {
			http.Error(w, "Error occurred while parsing form", http.StatusInternalServerError)
			return
		}

		Username := r.FormValue("username")
		Password := r.FormValue("password")
		DOB := r.FormValue("date")

		fmt.Fprintf(w, "Here are the details you entered for registration with us :- \n")
		fmt.Fprintf(w, "Username is :- %s", Username)
		fmt.Fprintf(w, "Password is :- %s", Password)
		fmt.Fprintf(w, "DOB is :- %s ", DOB)

	})

	http.ListenAndServe(":8080", nil)
}
