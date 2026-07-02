package handlers

import (
	"encoding/json"
	"hpc-backend/internal/models"
	"log"
	"net/http"
	"os"
	//"github.com/gorilla/mux"
)

func GetResults(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// 1. Leer el archivo JSON externo
	fileData, err := os.ReadFile("results/resultados_go.json")
	if err != nil {
		log.Printf("Error leyendo el archivo: %v", err)
		http.Error(w, "Error interno del servidor al cargar datos", http.StatusInternalServerError)
		return
	}

	// 2. Deserializar (Unmarshal) el JSON a nuestras estructuras de Go
	var results []models.Result
	if err := json.Unmarshal(fileData, &results); err != nil {
		log.Printf("Error decodificando el JSON: %v", err)
		http.Error(w, "Error procesando las métricas del clúster", http.StatusInternalServerError)
		return
	}

	// 3. Preparar la respuesta HTTP
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	// 4. Serializar (Marshal) las estructuras de vuelta a JSON para el cliente
	if err := json.NewEncoder(w).Encode(results); err != nil {
		log.Printf("Error enviando la respuesta: %v", err)
	}
}
