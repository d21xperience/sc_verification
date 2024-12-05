package model

import "gorm.io/gorm"

type Tenant struct {
	gorm.Model
	// ID   uint   `gorm:"primaryKey;autoIncrement" json:"tenant_id"`
	Name string `gorm:"not null"`
}