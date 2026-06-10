package router

import (
	"hpc-backend/internal/handlers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// API Endpoints
	r.HandleFunc("/api/results", handlers.GetResults).Methods("GET")

	return r
}
