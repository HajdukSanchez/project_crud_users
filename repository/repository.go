package repository

import (
	"context"

	"github.com/HajdukSanchez/project_crud_users/models"
)

// Repository for handle user process
type Repository interface {
	CreateUser(ctx context.Context, user *models.User) error
	ReadUser(ctx context.Context, id string) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id string) error
}

// Implementation for this abstract interface
var implementation Repository

// Function to handle dependency injection for this repository abstraction
func SetRepository(repository Repository) {
	implementation = repository
}

// Function handle by the abstraction
func CreateUser(ctx context.Context, user *models.User) error {
	return implementation.CreateUser(ctx, user)
}

// Function handle by the abstraction
func ReadUser(ctx context.Context, id string) (*models.User, error) {
	return implementation.ReadUser(ctx, id)
}

// Function handle by the abstraction
func UpdateUser(ctx context.Context, user *models.User) error {
	return implementation.UpdateUser(ctx, user)
}

// Function handle by the abstraction
func DeleteUser(ctx context.Context, id string) error {
	return implementation.DeleteUser(ctx, id)
}
