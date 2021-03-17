package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	fmt.Println("Web Server running at http://localhost:8080")

	tmpl := template.Must(template.ParseFiles("main.html"))

	// Registering request handler to default HTTP Server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "TODO list :",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: false},
			},
		}
		tmpl.Execute(w, data)
	})

	// An HTTP server has to listen on port to pass connections on to request handler
	http.ListenAndServe(":8080", nil)
}
