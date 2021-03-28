package repository

import (
	"github.com/garhus2020/blade/pkg/domain"
	"gorm.io/gorm"
)

// BookmarkRepository stores bookmarks
type BookmarkRepository struct {
	db *gorm.DB
}

// NewBookmarkRepository builds new bookmark repository using a db connection
func NewBookmarkRepository(db *gorm.DB) *BookmarkRepository {
	return &BookmarkRepository{
		db: db,
	}
}

// Create creates bookmark
func (r *BookmarkRepository) Create(bookmark *domain.Bookmark) (*domain.Bookmark, error) {
	r.db.AutoMigrate(&Bookmark{})
	r.db.Create(bookmark)
	return bookmark, nil
}

// GetAll gets all bookmarks
func (r *BookmarkRepository) GetAll() ([]*domain.Bookmark, error) {
	r.db.AutoMigrate(&Bookmark{})
	bookmarks := r.db.Find(&Bookmark)

	return bookmarks, nil
}
