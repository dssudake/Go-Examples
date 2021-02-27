package main

/*
	We need webserver to deploy web application which provides environment for web applications to run.
	In Go that environment is provided by net/http package and is compiled together
	with rest of your code to create a readily deployable standalone web app
	The net/http package from standard library contains all functionalities about HTTP protocol.
*/
import (
	"fmt"
	"net/http"
)

/*
	Request Handler function -
	 - receives all incomming HTTP connections from browsers, HTTP clients or API requests

	function receives two parameters -->
	 http.ResponseWriter interface- write your text/html response.
	 http.Request struct - contains all information about this HTTP request including things like URL or header fields
*/
func handler(writer http.ResponseWriter, request *http.Request) {
	// Setting the Content-Type header
	writer.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(writer, "<h1>Hello, you've requested : %s</h1>", request.URL.Path)
}

func main() {
	fmt.Println("Web Server running at http://localhost:8080")

	// Registering request handler to default HTTP Server
	http.HandleFunc("/", handler)

	// An HTTP server has to listen on port to pass connections on to request handler
	http.ListenAndServe(":8080", nil)
}
