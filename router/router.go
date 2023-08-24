package router

import (
	"auva/handlers"
	"auva/services"
	"net/http"

	"github.com/gorilla/mux"
)

func InitServerRouter(dependencies services.ServerDependencies) http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/", handlers.MainPageHandler()).Methods(http.MethodGet)
	router.HandleFunc("/submitted.html", handlers.SubmittedPageHandler()).Methods(http.MethodGet)
	router.HandleFunc("/auth/register", handlers.RegisterUser(dependencies.AuthService)).Methods(http.MethodPost)
	return router
}
