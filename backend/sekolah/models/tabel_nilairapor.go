package models

import "github.com/google/uuid"

type TabelNilaiAkhir struct {
	IDNilaiAkhir    uuid.UUID `gorm:"type:uuid;primary_key"`
	AnggotaRombelID uuid.UUID `gorm:"type:uuid"`
	MataPelajaranID int
	SemesterID      string
	NilaiPeng       float64
	PredikatPeng    string
	NilaiKet        float64
	PredikatKet     string
	NilaiSik        float64
	PredikatSik     string
	NilaiSiksos     float64
	PredikatSiksos  string
	PesertaDidikID  uuid.UUID `gorm:"type:uuid"`
	IDMinat         string
	Semester        int
}

// func MigrateDatabase() {
// 	DB.AutoMigrate(&TabelNilaiAkhir{})
// }
