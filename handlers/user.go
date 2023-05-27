package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/HajdukSanchez/project_crud_users/models"
	"github.com/HajdukSanchez/project_crud_users/repository"
	"github.com/HajdukSanchez/project_crud_users/server"
	"github.com/gorilla/mux"
)

// Response for create user - C
type CreateUserResponse struct {
	UserCreated bool `json:"user_created"`
}

// Response for update user - U
type UpdateUserResponse struct {
	UserUpdated bool `json:"user_updated"`
}

// Response for delete user - U
type DeleteUserResponse struct {
	DeleteUpdated bool `json:"user_deleted"`
}

// C - Create user
func CreateUserHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestUser = models.User{}
		err := json.NewDecoder(r.Body).Decode(&requestUser) // Try to decode request body into struct
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest) // Bad request from client
			return
		}

		err = repository.CreateUser(r.Context(), &requestUser)
		// Error creating user
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Response user created
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(CreateUserResponse{
			UserCreated: true,
		})
	}
}

// RUD user handler - Read, Update, Delete
func RUDUserHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			readUser(w, r)
		case http.MethodPut:
			updateUser(w, r)
		case http.MethodDelete:
			deleteUser(w, r)
		default:
			http.Error(w, "No method allowed", http.StatusInternalServerError)
			return
		}
	}
}

// R - Read user
func readUser(w http.ResponseWriter, r *http.Request) {
	userId, err := getIdParam(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := repository.ReadUser(r.Context(), userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// U - Update user
func updateUser(w http.ResponseWriter, r *http.Request) {
	var requestUser = models.User{}
	err := json.NewDecoder(r.Body).Decode(&requestUser) // Try to decode request body into struct
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // Bad request from client
		return
	}

	err = repository.UpdateUser(r.Context(), &requestUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(UpdateUserResponse{
		UserUpdated: true,
	})
}

// D - Delete user
func deleteUser(w http.ResponseWriter, r *http.Request) {
	userId, err := getIdParam(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = repository.DeleteUser(r.Context(), userId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(DeleteUserResponse{
		DeleteUpdated: true,
	})
}

// Get ID from path parameters
func getIdParam(r *http.Request) (string, error) {
	params := mux.Vars(r) // Get Path parameters to get ID of post like 'post/:ID'
	id := params["id"]
	return id, nil
}
