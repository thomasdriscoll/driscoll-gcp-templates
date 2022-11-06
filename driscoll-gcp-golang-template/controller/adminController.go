package controller

import (
	"io"
	"net/http"

	"github.com/thomasdriscoll/driscoll-gcp-templates/driscoll-gcp-golang-template/enums"
)

// Creates mux that ties into main app
func CreateAdminController() *http.ServeMux {
	mux := http.NewServeMux()
	controller := AdminControllerImpl{}

	// Create all endpoints for controller
	mux.HandleFunc(enums.MONITORING_URL, controller.Monitoring)
	mux.HandleFunc(enums.PING_URL, controller.Ping)

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
	io.WriteString(w, enums.PING_MESSAGE)
}
