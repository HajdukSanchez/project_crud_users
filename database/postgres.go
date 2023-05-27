package database

import (
	"context"
	"database/sql"
	"log"

	"github.com/HajdukSanchez/project_crud_users/models"
	_ "github.com/lib/pq" // Import library without use
)

// This repository will be work as a concrete implementation of user repository
type PostgresRepository struct {
	db *sql.DB
}

// Constructor
func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url) // Open SQL connection
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db}, nil
}

// repository implementation
func (repo *PostgresRepository) CreateUser(ctx context.Context, user *models.User) error {
	// SQL handle query
	_, err := repo.db.ExecContext(ctx, "INSERT INTO users (id, name, last_name, document, document_type) VALUES ($1, $2, $3, $4, $5)", user.Id, user.Name, user.LastName, user.Document, user.DocumentType)
	return err
}

// repository implementation
func (repo *PostgresRepository) ReadUser(ctx context.Context, id string) (*models.User, error) {
	// Query context return rows of data
	rows, _ := repo.db.QueryContext(ctx, "SELECT id, name, last_name, document, document_type FROM users WHERE id = $1", id)

	// Close database connection at the end
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// If there is some error getting data from database
	if err := rows.Err(); err != nil {
		return nil, err
	}

	var user = models.User{}
	for rows.Next() {
		// Try to map values from rows into model
		if err := rows.Scan(&user.Id, &user.Name, &user.LastName, &user.Document, &user.DocumentType); err == nil {
			return &user, nil // Everything ok
		}
	}

	return &user, nil
}

// repository implementation
func (repo *PostgresRepository) UpdateUser(ctx context.Context, user *models.User) error {
	// SQL handle query
	_, err := repo.db.ExecContext(ctx, "UPDATE users SET id = $1, name = $2, last_name = $3, document = $4, document_type = $5 WHERE id = $1", user.Id, user.Name, user.LastName, user.Document, user.DocumentType)

	return err
}

// repository implementation
func (repo *PostgresRepository) DeleteUser(ctx context.Context, id string) error {
	// SQL handle query
	_, err := repo.db.ExecContext(ctx, "DELETE FROM users WHERE id = $1", id)

	return err
}
