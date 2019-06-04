package models

// Todo is the model for graphql
type Todo struct {
	ID     string
	Text   string
	Done   bool
	UserID string
}
