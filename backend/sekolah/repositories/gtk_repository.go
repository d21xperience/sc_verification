package repositories

import (
	"sekolah/models"

	"gorm.io/gorm"
)

type GTKRepository interface {
	FindAll() ([]models.GTKAPI, error)
	// Create(gtk models.GTKAPI) (*models.GTKAPI, error)
}

type gtkRepository struct {
	DB *gorm.DB
}

func NewGTKRepository(DB *gorm.DB) GTKRepository {
	return &gtkRepository{DB}
}

func (g *gtkRepository) FindAll() ([]models.GTKAPI, error) {
	var gtk []models.GTKAPI
	err := g.DB.Find(&gtk).Error
	return gtk, err
}

// func (g *gtkRepository) Create(gtk models.GTKAPI) (*models.GTKAPI, error) {

// }
