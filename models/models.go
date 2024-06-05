package models

import (
	"go-crud/models/user"
)

// GetModels returns a slice of all models
func GetModels() []interface{} {
	return []interface{}{
		&user.User{},
		&user.LoginDetails{}, // Add LoginDetails model
		// Add other models here if needed
	}
}
