package database

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/firebase007/go-rest-api-with-fiber/config"
	"github.com/go-sql-driver/mysql"
)

// Database instance
var DB *sql.DB

// Connect function
func Connect() error {
	var err error

	p := config.Config("DB_PORT")
	fmt.Printf("DB_PORT: %s\n", p)

	cfg := mysql.Config{
		User:   config.Config("DB_USER"),
		Passwd: config.Config("DB_PASSWORD"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "auth",
	}

	// because our config function returns a string, we are parsing our str to int here - hack ??
	// port, err := strconv.ParseUint(p, 10, 32)

	// if err != nil {
	// 	fmt.Println("Error parsing str to int")
	// }
	// fmt.Printf("port: %d\n", port)

	// DB, err = sql.Open("sql", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME")))
	DB, err = sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return errors.New("ping database error")
	}

	CreateProductTable()
	return nil
}
