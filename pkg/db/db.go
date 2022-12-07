package db

import (
	"log"

	"github.com/DavidKorochik/issues4missions-backend/pkg/config"
	"github.com/DavidKorochik/issues4missions-backend/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func ConnectToDB(config config.Config) *Database {
	db, err := gorm.Open(postgres.Open(config.DBUrl), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot connect to the database ", err)
		return &Database{
			DB: db,
		}
	}

	autoMigrateModels(db)

	return &Database{
		DB: db,
	}
}

func autoMigrateModels(db *gorm.DB) {
	db.AutoMigrate(&models.Issue{}, &models.User{}, &models.Department{}, &models.Role{})
}
