package user

import (
	"time"

	"gorm.io/gorm"
)

// User represents the user model
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100;not null;unique"`
	Email     string `gorm:"size:255;not null;unique"`
	Password  string `gorm:"size:255;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"` // For soft delete
}
