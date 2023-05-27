package models

// User model
type User struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	LastName     string `json:"last_name"`
	Document     string `json:"document"`
	DocumentType string `json:"document_type"`
}
