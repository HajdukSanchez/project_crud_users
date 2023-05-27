package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/HajdukSanchez/project_crud_users/server"
)

// Response to return to health status
type HealthResponse struct {
	Status string `json:"status"`
}

func HealthHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json") // Specify the response is a json response
		w.WriteHeader(http.StatusOK)                       // 200

		json.NewEncoder(w).Encode(HealthResponse{
			Status: "Ok",
		}) // Encoder to create new response
	}
}
