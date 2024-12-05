package repositories

import "gorm.io/gorm"

import "sekolah/models"

type SekolahRepository interface {
	Save(sekolah models.Sekolah)
	Update(sekolah models.Sekolah)
	
}

type SekolahRepositoryIml struct {
	DB *gorm.DB
}