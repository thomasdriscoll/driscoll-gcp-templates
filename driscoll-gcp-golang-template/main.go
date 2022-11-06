package main

import (
	"log"
	"net/http"

	controller "github.com/thomasdriscoll/driscoll-gcp-templates/driscoll-gcp-golang-template"
)

func setup() *http.ServeMux {
	// Create controllers
	adminMux := controller.CreateAdminController()
	// add
	mux := http.NewServeMux()
	mux.Handle("/admin/", http.StripPrefix("/admin", adminMux))
	return mux
}

func main() {
	mux := setup()
	log.Printf("Starting http server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", mux))
}
