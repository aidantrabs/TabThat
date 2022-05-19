package services

import (
	"example/bookmark-api/models"
)

type BookmarkService interface {
	CreateBM (*models.Bookmark) error 
	GetBM(*string) (*models.Bookmark, error)
	GetAllBM() ([]*models.Bookmark, error)
	UpdateBM(*models.Bookmark) error
	DeleteBM(*string) error
}