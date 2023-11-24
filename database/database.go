package database

import (
	"errors"
	"fmt"

	"github.com/firebase007/go-rest-api-with-fiber/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database instance
var DB *gorm.DB

// Connect function
func Connect() error {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		"127.0.0.1",
		"3306",
		"auth",
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

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

	CreateProductTable()
	return nil
}
