package model

import (
	"time"
)

type User struct {
	// gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	// Role           string `gorm:"not null"`
	Nama        string `gorm:"size:100" json:"nama"`
	JK          string `gorm:"size:100" json:"jk"`
	Phone       string `gorm:"size:100" json:"phone"`
	TptLahir    string `gorm:"size:100" json:"tpt_lahir"`
	TglLahir    time.Time
	AlamatJalan string    `gorm:"size:100" json:"alamat_jalan"`
	KotaKab     string    `gorm:"size:100" json:"kota_kab"`
	Prov        string    `gorm:"size:100" json:"prov"`
	KodePos     string    `gorm:"size:100" json:"kode_pos"`
	NamaAyah    string    `gorm:"size:100" json:"nama_ayah"`
	NamaIbu     string    `gorm:"size:100" json:"nama_ibu"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	RoleID         uint   `gorm:"not"`
	TenantID       uint   `gorm:"foreignKey:ID"`
	ProfilePicture string `gorm:"type:varchar(255)" json:"photo_url"`
}
