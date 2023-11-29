package database

import (
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database instance
var DB *gorm.DB

// Connect function
func Connect() error {
	var err error

	dsn := "postgresql://root@db:26257/defaultdb?sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	dbObject, err := DB.DB()
	if err != nil {
		return err
	}

	if err = dbObject.Ping(); err != nil {
		return errors.New("ping database error")
	}

	CreateDatabase()
	CreateProductTable()
	return nil
}
