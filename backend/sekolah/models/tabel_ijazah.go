package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Ijazah struct {
	gorm.Model
	ID             string    `gorm:"primaryKey"`
	PesertaDidikID uuid.UUID `gorm:"foreignkey"`
	RerataNilai    float32   `gorm:"reratan_nilai"`
	NomorIjazah    string    `gorm:"nomor_ijazah"`
}
