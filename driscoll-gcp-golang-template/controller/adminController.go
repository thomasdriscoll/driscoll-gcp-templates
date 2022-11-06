package controller

import (
	"io"
	"net/http"
)

// Creates mux that ties into main app
func CreateAdminController() *http.ServeMux {
	mux := http.NewServeMux()
	controller := AdminControllerImpl{}

	// Create all endpoints for controller
	mux.HandleFunc("/monitoring", controller.Monitoring)
	mux.HandleFunc("/ping", controller.Ping)

	return mux
}

/* Start of AdminController logic */
type AdminController interface {
	Monitoring(w http.ResponseWriter, r *http.Request)
	Ping(w http.ResponseWriter, r *http.Request)
}

type AdminControllerImpl struct {
}

func (ac AdminControllerImpl) Monitoring(w http.ResponseWriter, r *http.Request) {

}

func (ac AdminControllerImpl) Ping(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, string("I am alive"))
}
