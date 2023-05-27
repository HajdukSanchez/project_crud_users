package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/HajdukSanchez/project_crud_users/handlers"
	"github.com/HajdukSanchez/project_crud_users/server"
	"github.com/HajdukSanchez/project_crud_users/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env") // Load the environments

	if err != nil {
		log.Fatal("Error loading environments")
	}

	// Get all environments
	PORT := os.Getenv("PORT")
	DATA_BASE_URL := os.Getenv("DATA_BASE_URL")

	// Create the new server
	server, err := server.NewServer(context.Background(), &server.Config{
		Port:  PORT,
		DBUrl: DATA_BASE_URL,
	})

	if err != nil {
		log.Fatal("Error creating server")
	}

	server.Start(BindRoutes) // Start the server
}

// Function to handle routes and start server
func BindRoutes(server server.Server, router *mux.Router) {
	// Define endpoints and methods for endpoints
	router.HandleFunc(utils.Status, handlers.HealthHandler(server)).Methods(http.MethodGet)
	router.HandleFunc(utils.NewUser, handlers.CreateUserHandler(server)).Methods(http.MethodPost)
	router.HandleFunc(utils.RUDUser, handlers.RUDUserHandler(server)).Methods(http.MethodGet, http.MethodPut, http.MethodDelete)
}
