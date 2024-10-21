// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		http.ServeFile(w, r, "./index.html")
// 	})

// 	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
// 		if r.Method != http.MethodPost {
// 			http.Error(w, "Error occured, not the right http method :-", http.StatusMethodNotAllowed)
// 			return
// 		}

// 		if err := r.ParseForm(); err != nil {
// 			http.Error(w, "error parsing the form data :- ", http.StatusInternalServerError)
// 			return
// 		}

// 		Username := r.FormValue("username")
// 		Password := r.FormValue("password")
// 		Email := r.FormValue("email")

// 		fmt.Fprintf(w, "YOUR RESPONSE HAS BEEN RECORDED SUCCESSFULLY. HERE ARE THE DETAILS YOU ENTERED :-\n")
// 		fmt.Fprintf(w, "Username :- %s\n", Username)
// 		fmt.Fprintf(w, "Password :- %s\n", Password)
// 		fmt.Fprintf(w, "Email :- %s\n", Email)
// 	})

// 	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {

// 	})
// 	var port string = ":8080"

// 	fmt.Printf("Server started at the port :-  %s", port)

// 	if err := http.ListenAndServe(port, nil); err != nil {
// 		log.Fatal("Error starting the server at port")
// 	}

// }

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	fmt.Fprintf(w, "User ID is :- %s", userID)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // Parse form data
	name := r.FormValue("name")
	fmt.Fprintf(w, "New user created: %s", name)
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/users/{id}", GetUserHandler).Methods("GET")
	r.HandleFunc("/users", createUserHandler).Methods("POST")

	log.Println("Server Started at port : 8080")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Error starting server")
	}

}
