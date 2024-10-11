package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the root path of my server !")
	})
	var port string = ":8080"
	fmt.Println("Server is up and running at :- ", port)
	http.ListenAndServe(port, nil)

}
