package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 54533
	user     = "postgres"
	password = ""
	dbName   = "postgres"
)

var DB *gorm.DB

func InitDatabase() {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	var err error
	DB, err = gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatal("Terjadi error dg error sbb: ", err)
	}
	// DB.AutoMigrate(&model.User{})
	// SeedProfiles(DB, 5)
	log.Println("Database connected successfully!")
}
