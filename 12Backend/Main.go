package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "WELCOME TO OUR SERVER BOSS !")
	})

	// Agar main ek static file serve karna chahta hu in golang then there are bsically 2 steps =>

	// step 1

	fs := http.FileServer(http.Dir("./Static"))

	// jo directory mein se serve karna hai usse fs mein store karlia using http.FileServer

	// step 2

	http.Handle("/form/", http.StripPrefix("/form/", fs))

	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Error occurred :-", http.StatusMethodNotAllowed)
		}

		if err := r.ParseForm(); err != nil {
			http.Error(w, "an error has occurred", http.StatusInternalServerError)
		}

		Username := r.FormValue("username")
		Password := r.FormValue("password")
		Telephone := r.FormValue("telephone")

		fmt.Fprintf(w, "Form has been submitted successfully with the following inputs :- \n")
		fmt.Fprintf(w, "Username is :- %s\n", Username)
		fmt.Fprintf(w, "Password is :- %s\n", Password)
		fmt.Fprintf(w, "Telephone is :- %s\n", Telephone)

	})

	fmt.Println("Server Successfully started at port : 8080")

	http.ListenAndServe(":8080", nil)

}
