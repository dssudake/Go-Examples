package main

import (
	"fmt"
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	// Setting the Content-Type header
	w.Header().Set("Content-Type", "text/html")

	fmt.Fprintf(w, "<center style=\"margin-top:300px\">")

	// Writing Response according to the routes
	if r.URL.Path == "/" || r.URL.Path == "/home" {
		fmt.Fprintf(w, "<h1>Welcome to Go Webapp :)</h1><p>This is Home Page</p>")
	} else if r.URL.Path == "/about" {
		fmt.Fprintf(w, "<h1>This is About Page</h1><p>Welcome to Go Webapp :)</p>")
	} else {
		fmt.Fprintf(w, "<h1>We could not find the page you were looking for :(</h1><p>Please email us if you keep being send to an invalid page</p>")
	}

	fmt.Fprintf(w, "</center>")
}

func main() {
	fmt.Println("Web Server running at http://localhost:8080")

	/*
		ServeMux is an HTTP request multiplexer.
		It matches URL of each incoming request against list of registered patterns
		and calls the handler for the pattern that most closely matches the URL.
	*/
	mux := &http.ServeMux{}

	// Registering request handler
	mux.HandleFunc("/", handleFunc)

	// An HTTP server has to listen on port to pass connections on to request handler
	http.ListenAndServe(":8080", mux)
}
