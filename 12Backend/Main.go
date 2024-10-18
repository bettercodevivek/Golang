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
	Dob      string `json:"dob"`
}

var mu sync.Mutex
var userFile = "users.json"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "WELCOME TO VIVEK'S SERVER !!!")
	})
	// root path of the server

	fs := http.FileServer(http.Dir("./Static"))
	http.Handle("/form/", http.StripPrefix("/form/", fs))

	// Handle form submission
	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "error occurred", http.StatusMethodNotAllowed)
			return
		}

		if err := r.ParseForm(); err != nil {
			http.Error(w, "error occurred while parsing form", http.StatusInternalServerError)
			return
		}

		Username := r.FormValue("username")
		Password := r.FormValue("password")
		DOB := r.FormValue("date")

		fmt.Fprintf(w, "NEW USER REGISTERED WITH THE FOLLOWING DETAILS :-\n")
		fmt.Fprintf(w, "Username is :-%s\n", Username)
		fmt.Fprintf(w, "Password is :-%s\n", Password)
		fmt.Fprintf(w, "Date Of Birth is :-%s\n", DOB)

		newUser := user{Username: Username, Password: Password, Dob: DOB}

		if err := SaveUserToJson(newUser); err != nil {
			http.Error(w, "error occurred in saving to json: "+err.Error(), http.StatusInternalServerError)
			return
		}
	})

	fmt.Println("Server started on port 8000")
	http.ListenAndServe(":8000", nil)
}

func SaveUserToJson(newUser user) error {
	var users []user

	mu.Lock()
	defer mu.Unlock()

	// Check if the file exists
	if _, err := os.Stat(userFile); err == nil {
		// File exists, open it and decode existing users
		file, err := os.OpenFile(userFile, os.O_RDWR, 0644)
		if err != nil {
			return err
		}
		defer file.Close()

		// Decode existing users if file is not empty
		if err := json.NewDecoder(file).Decode(&users); err != nil && err.Error() != "EOF" {
			return err
		}
	} else if os.IsNotExist(err) {
		// File does not exist, create an empty list
		users = []user{}
	} else {
		// Some other error
		return err
	}

	// Append the new user to the list
	users = append(users, newUser)

	// Reopen the file in write mode, truncating it so we overwrite the content
	file, err := os.OpenFile(userFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode the updated users list to the file
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(users); err != nil {
		return err
	}

	return nil
}
