package main

import (
	"fmt"
	"net/http"

	"groupieTracker/features"
)
// main is the entry point of the application. It prints a startup message,
// sets up the routes for the web server, and starts the HTTP server on port 8080.
// If the server fails to start or encounters an error while running, it prints
// an error message to the console.
func main() {
	fmt.Println("localhost is lunched : http://localhost:8080")
	features.SetupRoutes()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("main error : %v", err)
	}
}