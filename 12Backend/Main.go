package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
	DOB      string `json:"DOB"`
}

var mu sync.Mutex

var userfile = "users.json"

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
			return
		}

		if err := r.ParseForm(); err != nil {
			http.Error(w, "an error has occurred", http.StatusInternalServerError)
			return
		}

		Username := r.FormValue("username")
		Password := r.FormValue("password")
		Date := r.FormValue("date")

		fmt.Fprintf(w, "Form has been submitted successfully with the following inputs :- \n")
		fmt.Fprintf(w, "Username is :- %s\n", Username)
		fmt.Fprintf(w, "Password is :- %s\n", Password)
		fmt.Fprintf(w, "DOB is :- %s\n", Date)

		newUser := user{Username, Password, Date}

		if err := SaveUsertoJSON(newUser); err != nil {
			http.Error(w, "An error occurred while saving a new user", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "User data saved to JSON file!\n")
	})

	fmt.Println("Server Successfully started at port : 8080")

	http.ListenAndServe(":8080", nil)

}

func SaveUsertoJSON(newUser user) error {
	mu.Lock()
	defer mu.Unlock()

	var users []user

	if _, err := os.Stat(userfile); err != nil {
		file, err := os.OpenFile(userfile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			return err
		}
		defer file.Close()

		if err := json.NewDecoder(file).Decode(&users); err != nil {
			return err
		}
	}

	users = append(users, newUser)

	file, err := os.Create(userfile)
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(users); err != nil {
		return err
	}
	return nil
}
