package features

import "net/http"

// SetupRoutes configures the HTTP server's routing. It sets up a file server
// to serve static files from the "static" directory, and registers handler
// functions for the root path ("/").
// The root path is handled by the main Handler function.
func SetupRoutes() {
	http.HandleFunc("/", Handler)
}
