package main

import (
	"fmt"
	"net/http"
)

// Declare a handler which writes a plain-text response with information about the
// application status, operating environment and version.
func (app *application) indexHandler(w http.ResponseWriter, r *http.Request) {
	app.logger.Println("index called")
	fmt.Fprintln(w, "Accessed the index page, congrats ! :)")
}
