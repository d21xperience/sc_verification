package model

type Role struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"role_id"`
	Name string `gorm:"not null"`
}
