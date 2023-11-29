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

	//dsn := "postgresql://user:password@db:26257/database?sslmode=disable"

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
	// 	config.Config("DB_HOST"),
	// 	config.Config("DB_USER"),
	// 	config.Config("DB_PASSWORD"),
	// 	config.Config("DB_NAME"),
	// 	config.Config("DB_PORT"),
	// )

	dsn := "postgresql://gorm:gorm@roach:26257/products?sslmode=disable"

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
