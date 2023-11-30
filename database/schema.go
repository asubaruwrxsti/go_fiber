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
			category TEXT
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
		CREATE DATABASE IF NOT EXISTS products;
		`,
	)

	return nil
}
