package handler

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/firebase007/go-rest-api-with-fiber/model"

	"github.com/firebase007/go-rest-api-with-fiber/database"
)

// GetAllProducts godoc
// @Summary Get all products
// @Description Get all products
// @Tags products
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api [get]
func GetAllProducts(c *fiber.Ctx) error {

	// query product table in the database
	rows, err := database.DB.Find(&model.Product{}).Rows()
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
		return err
	}

	defer rows.Close()

	result := model.Products{}

	for rows.Next() {
		product := model.Product{}
		err := rows.Scan(&product.Name, &product.Description, &product.Category, &product.Amount)
		// Exit if we get an error
		if err != nil {
			c.Status(500).JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
			return err
		}

		// Append Product to Products
		result.Products = append(result.Products, product)
	}

	// Return Products in JSON format
	if err := c.JSON(&fiber.Map{
		"success": true,
		"product": result,
		"message": "All product returned successfully",
	}); err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return err
	}
	return nil
}

// GetSingleProduct godoc
// @Summary Get a product
// @Description Get a product
// @Tags products
// @Accept */*
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/{id} [get]
func GetSingleProduct(c *fiber.Ctx) error {

	id := c.Params("id")
	product := model.Product{}

	// query product database
	row, err := database.DB.Take(&product, id).Rows()
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return err
	}

	defer row.Close()

	// iterate through the values of the row
	for row.Next() {
		switch err := row.Scan(&id, &product.Amount, &product.Name, &product.Description, &product.Category); err {
		case sql.ErrNoRows:
			log.Println("No rows were returned!")
			c.Status(500).JSON(&fiber.Map{
				"success": false,
				"message": err,
			})
		case nil:
			log.Println(product.Name, product.Description, product.Category, product.Amount)
		default:
			//   panic(err)
			c.Status(500).JSON(&fiber.Map{
				"success": false,
				"message": err,
			})
		}

	}

	// return product in JSON format
	if err := c.JSON(&fiber.Map{
		"success": false,
		"message": "Successfully fetched product",
		"product": product,
	}); err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return err
	}
	return nil
}

// CreateProduct godoc
// @Summary Create a product
// @Description Create a product
// @Tags products
// @Accept */*
// @Produce json
// @Param name body string true "Product Name"
// @Param description body string true "Product Description"
// @Param category body string true "Product Category"
// @Param amount body int true "Product Amount"
// @Success 200 {object} map[string]interface{}
// @Router /api [post]
func CreateProduct(c *fiber.Ctx) error {

	// Instantiate new Product struct
	p := new(model.Product)

	//  Parse body into product struct
	if err := c.BodyParser(p); err != nil {
		log.Println(err)
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return err
	}

	// Insert Product into database
	res, err := database.DB.Create(&p).Rows()
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return err
	}
	// Print result
	log.Println(res)

	// Return Product in JSON format
	if err := c.JSON(&fiber.Map{
		"success": true,
		"message": "Product successfully created",
		"product": p,
	}); err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Error creating product",
		})
		return err
	}
	return nil
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Delete a product
// @Tags products
// @Accept */*
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/{id} [delete]
func DeleteProduct(c *fiber.Ctx) error {

	id := c.Params("id")

	// query product table in database
	res, err := database.DB.Delete(&model.Product{}, id).Rows()
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
		return err
	}

	// Print result
	log.Println(res)

	// return product in JSON format
	if err := c.JSON(&fiber.Map{
		"success": true,
		"message": "product deleted successfully",
	}); err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
		return err
	}
	return nil
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags health
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/health [get]
func Ping(c *fiber.Ctx) error {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	if err := c.JSON(res); err != nil {
		return err
	}

	return nil
}
