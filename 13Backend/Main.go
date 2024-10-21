package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html")
	})

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Error occured, not the right http method :-", http.StatusMethodNotAllowed)
			return
		}

		if err := r.ParseForm(); err != nil {
			http.Error(w, "error parsing the form data :- ", http.StatusInternalServerError)
			return
		}

		Username := r.FormValue("username")
		Password := r.FormValue("password")
		Email := r.FormValue("email")

		fmt.Fprintf(w, "YOUR RESPONSE HAS BEEN RECORDED SUCCESSFULLY. HERE ARE THE DETAILS YOU ENTERED :-\n")
		fmt.Fprintf(w, "Username :- %s\n", Username)
		fmt.Fprintf(w, "Password :- %s\n", Password)
		fmt.Fprintf(w, "Email :- %s\n", Email)
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {

	})
	var port string = ":8080"

	fmt.Printf("Server started at the port :-  %s", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Error starting the server at port")
	}

}
