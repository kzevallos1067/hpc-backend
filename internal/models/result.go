package models

//Resultados generales
type Result struct {
	Instance_Nodes    string  `json:"instance_nodes"`
	Minimal_distante  float64 `json:"minimal_distance"`
	Sequence          Path    `json:"secuence"`
	Ts                float64 `json:"Ts"`
	Tp                float64 `json:"Tp"`
	Energy            float64 `json:"energy"`
	Processors        int64   `json:"processors"`
	Eficency_Score    float64 `json:"eficency_score"`
	SpeedUp_Score     float64 `json:"speedup_score"`
	Consumption_Score float64 `json:"consumption_score"`
}

type Path struct {
	Sequence []NodeResult `json:"sequence"`
	Total    float64      `json:"total_distance"` // Distancia total del recorrido
}

type NodeResult struct {
	ID       int64   `json:"id"`       // Identificador único
	Distance float64 `json:"distance"` // Distancia acumulada o parcial
}
