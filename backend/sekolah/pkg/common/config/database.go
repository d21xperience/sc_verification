package config

import (
	"fmt"
	"sekolah/helper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DbInfo struct {
	host     string
	port     uint16
	dbName   string
	user     string
	password string
}

func DatabaseConnection() *gorm.DB {
	c, err := LoadConfig()
	helper.ErrorPanic(err)
	var dbInfo = DbInfo{c.DBUrl, c.DBPort, c.DBName, c.DBUser, c.DBPassword}

	sqlInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", dbInfo.host, dbInfo.user, dbInfo.password, dbInfo.dbName, dbInfo.port)
	fmt.Println(sqlInfo)
	// sqlInfo := "host=localhost user=postgres dbname=dbraporsp port=57152 sslmode=disable"
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	helper.ErrorPanic(err)

	return db
}
