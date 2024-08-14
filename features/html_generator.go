package features

import (
	"html/template"
	"net/http"
)

// OpenHtml parses an HTML template file, executes it with the provided data,
// and writes the result to the HTTP response. If any errors occur during
// template parsing or execution, it calls the ErrorHandler function with
// a 500 Internal Server Error status.
func OpenHtml(fileName string, response http.ResponseWriter, data any) {
	temp, err := template.ParseFiles("templates/" + fileName)
	if err != nil {
		ErrorHandler(response, http.StatusInternalServerError)
		return
	}
	err = temp.Execute(response, data)
	if err != nil {
		ErrorHandler(response, http.StatusInternalServerError)
		return
	}
}

// ErrorHandler manages HTTP error responses. It sets the appropriate status code
// and renders a specific error page template based on the error type:
// 404 for Not Found, 500 for Internal Server Error, and 405 for Method Not Allowed.
// It uses the OpenHtml function to render these error pages.
func ErrorHandler(response http.ResponseWriter, status int) {
	response.WriteHeader(status)
	if status == http.StatusNotFound {
		OpenHtml("404.html", response, nil)
	} else if status == http.StatusInternalServerError {
		OpenHtml("500.html", response, nil)
	} else if status == http.StatusMethodNotAllowed {
		OpenHtml("405.html", response, nil)
	}
}
