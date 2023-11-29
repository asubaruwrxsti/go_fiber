package database

// CreateProductTable ...
func CreateProductTable() error {
	dbObject, err := DB.DB()
	if err != nil {
		return err
	}

	dbObject.Query(`
		CREATE TABLE IF NOT EXISTS products (
			id SERIAL PRIMARY KEY,
			amount INTEGER,
			name TEXT UNIQUE,
			description TEXT,
			category TEXT NOT NULL
		);`,
	)
	return nil
}

func CreateDatabase() error {
	dbObject, err := DB.DB()
	if err != nil {
		return err
	}

	dbObject.Query(`
	CREATE USER gorm WITH PASSWORD 'gorm';
	GRANT ALL PRIVILEGES ON DATABASE gorm TO gorm;
	`,
	)

	dbObject.Query(`
		CREATE DATABASE IF NOT EXISTS products;
		GRANT ALL PRIVILEGES ON DATABASE gorm TO gorm;
		`,
	)

	return nil
}
