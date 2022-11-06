package main

import (
	"log"
	"net/http"

	"github.com/thomasdriscoll/driscoll-gcp-templates/driscoll-gcp-golang-template/controller"
	"github.com/thomasdriscoll/driscoll-gcp-templates/driscoll-gcp-golang-template/enums"
)

func setup() *http.ServeMux {
	// Create controllers
	adminMux := controller.CreateAdminController()
	// add
	mux := http.NewServeMux()
	mux.Handle(enums.ADMIN_URL+"/", http.StripPrefix(enums.ADMIN_URL, adminMux))
	return mux
}

func main() {
	mux := setup()
	log.Printf("Starting http server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", mux))
}
