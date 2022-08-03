package db

import (
	"os"

	"github.com/adnanbrq/slugify/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func IsConnected() bool {
	if DB == nil {
		return false
	}

	sql, ok := DB.DB()
	if ok != nil {
		return false
	}

	return sql.Ping() == nil
}

func Connect() error {
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = db
	if autoMigrate := os.Getenv("AUTO_MIGRATE") == "true"; autoMigrate {
		db.AutoMigrate(&entity.Link{})
	}
	return nil
}
