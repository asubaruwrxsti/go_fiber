package database

// CreateProductTable ...
func CreateProductTable() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}

	sqlDB.Query(`CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    amount integer,
    name text UNIQUE,
    description text,
    category text NOT NULL
)
`)
	return nil

}
