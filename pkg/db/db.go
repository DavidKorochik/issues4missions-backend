package db

import (
	"log"

	"github.com/DavidKorochik/issues4missions-backend/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB(config config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.DBUrl), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot connect to the database ", err)
		return nil
	}

	autoMigrateModels(db)

	return db
}

func autoMigrateModels(db *gorm.DB) {
	db.AutoMigrate()
}
