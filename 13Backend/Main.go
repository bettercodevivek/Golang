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

	var port string = ":8080"

	fmt.Printf("Server started at the port :-  %s", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Error starting the server at port")
	}

}
