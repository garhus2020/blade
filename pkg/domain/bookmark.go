package domain

import (
	"time"

	"gorm.io/gorm"
)

// Bookmark domain entity
type Bookmark struct {
	gorm.Model
	ID        int
	Name      string
	URI       string
	Category  string
	CreatedAt time.Time
}
