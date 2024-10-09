// package main

// // let us see a simple http server in golang using "net/http" and "log" packages.

// import (
// 	"fmt"
// 	"log"      // used to log any fatal errors when server is run
// 	"net/http" // main http package for running a http server
// )

// // func handler(w http.ResponseWriter, r *http.Request) {
// // 	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
// // }

// func main() {
// 	// This function associates a URL path (route) with a handler function. Here, it takes two arguments:
// 	//This is the route path. In this case, / is the root path of your server. Any request made to this server will match this route.
// 	// This is an inline function with two parameters (w http.ResponseWriter and r *http.Request). The first parameter w is used to send a response back to the client, and the second parameter r represents the client’s HTTP request.
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "HEY! WELCOME TO MY FIRST GOLANG SERVER !")
// 	})
// 	port := ":8000"
// 	log.Fatal(http.ListenAndServe(port, nil))

// 	// http.ListenAndServe: This function starts the HTTP server on the specified port (:8000 in this case) and starts listening for incoming HTTP requests.

// 	// The second parameter is nil, which tells the server to use the default HTTP request multiplexer (http.DefaultServeMux). Since you already associated a route with http.HandleFunc, this will automatically be registered with the default multiplexer.

// 	// log.Fatal(...): If there’s an error when starting the server (e.g., the port is already in use), log.Fatal will log the error and immediately stop the program
// }

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "HELLO THIS SERVER IS OWNED BY VIVEK")
	})

	http.ListenAndServe(":80", nil)
}
