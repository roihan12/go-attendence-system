package database

import (
	"fmt"

	"github.com/roihan12/technical-test-kalbe/app/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDBPostgres(app config.AppConfig) *gorm.DB {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", app.DBHost, app.DBUser, app.DBPass, app.DBName, app.DBPort)

	DB, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return DB
}