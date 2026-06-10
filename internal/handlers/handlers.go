package handlers

import (
	"encoding/json"
	"hpc-backend/internal/models"
	"net/http"
	//"github.com/gorilla/mux"
)

func GetResults(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Read dynamic URL path parameters via Gorilla Mux
	//params := mux.Vars(r)
	//instanceNodes := params["instance_nodes"]

	sequence := models.Path{
		Sequence: []models.NodeResult{
			{ID: 1, Distance: 0},    // nodo inicial
			{ID: 2, Distance: 12.5}, // distancia desde A a B
			{ID: 3, Distance: 7.3},  // distancia desde B a C
		},
		Total: 19.8, // suma de las distancias
	}
	result := models.Result{
		Instance_Nodes:   "5", //instanceNodes,
		Minimal_distante: sequence.Total,
		Sequence:         sequence,
		Ts:               1.5,
		Tp:               0.5,
		Energy:           2.45, //medido en kW
		Processors:       4,
	}
	json.NewEncoder(w).Encode(result)
}
