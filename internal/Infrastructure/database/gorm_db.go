package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewGormDB(debug bool) *gorm.DB {
	connStr := os.Getenv("DB")

	gormConfig := &gorm.Config{}
	if debug {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}

	var db *gorm.DB
	var err error

	for {
		db, err = gorm.Open(postgres.Open(connStr), gormConfig)
		if err != nil {
			fmt.Println("failed to connect database:", err)
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Token{},
		&models.ApiAuth{},
		&models.Log{},
	)
	if err != nil {
		log.Fatal("failed to migrate database:", err)
	}

	return db
}
