package main

import (
	"hpc-backend/internal/router"
	"log"
	"net/http"
)

func main() {
	r := router.SetupRouter()

	log.Println("Server executing on port :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
