package database

import (
	"errors"
	"fmt"

	"github.com/firebase007/go-rest-api-with-fiber/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database instance
var DB *gorm.DB

// Connect function
func Connect() error {
	var err error

	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		config.Config("DB_HOST"),
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_NAME"),
		config.Config("DB_PORT"),
		config.Config("DB_TIMEZONE"),
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}

	if err = sqlDB.Ping(); err != nil {
		return errors.New("ping database error")
	}

	CreateDatabase()
	CreateProductTable()
	return nil
}
