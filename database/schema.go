package database

// CreateProductTable ...
func CreateProductTable() error {
	dbObject, err := DB.DB()
	if err != nil {
		return err
	}

	dbObject.Query(`
		use test;
		`,
	)

	dbObject.Query(`
		CREATE TABLE IF NOT EXISTS products (
			id SERIAL PRIMARY KEY,
			name TEXT UNIQUE,
			description TEXT,
			category TEXT,
			amount INTEGER
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
		CREATE DATABASE IF NOT EXISTS test;
		`,
	)

	return nil
}
