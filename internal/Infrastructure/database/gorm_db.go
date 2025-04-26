package database

import (
	"log"

	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewGormDB(connStr string, debug bool) *gorm.DB {
	gormConfig := &gorm.Config{}
	if debug {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}

	db, err := gorm.Open(postgres.Open(connStr), gormConfig)
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Token{},
	)
	if err != nil {
		log.Fatal("failed to migrate database:", err)
	}

	return db
}
