package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/firebase007/go-rest-api-with-fiber/model"

	"github.com/firebase007/go-rest-api-with-fiber/database"
	"github.com/go-playground/validator/v10"
)

type (
	ErrorResponse struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface{}
	}

	XValidator struct {
		validator *validator.Validate
	}

	GlobalErrorHandlerResp struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
)

// This is the validator instance
// for more information see: https://github.com/go-playground/validator
var validate = validator.New()

func (v XValidator) Validate(data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem ErrorResponse

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

// Home godoc
// @Summary Hello World.
// @Description Hello World.
// @Tags home
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func Home(c *fiber.Ctx) error {
	return c.Render("home", fiber.Map{
		"Title":   "Hello, World!",
		"Message": "Hello, World!",
	})
}

// GetAllProducts godoc
// @Summary Get all products
// @Description Get all products
// @Tags products
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api [get]
func GetAllProducts(c *fiber.Ctx) error {
	// Query all products from the database
	var result model.Products
	if err := database.DB.Find(&result.Products).Error; err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}

	// Return Products in JSON format
	return c.JSON(&fiber.Map{
		"success": true,
		"product": result,
		"message": "All products returned successfully",
	})
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

	// Query single product from the database
	if err := database.DB.First(&product, id).Error; err != nil {
		return c.Status(404).JSON(&fiber.Map{
			"success": false,
			"message": "Product not found",
		})
	}

	// Return product in JSON format
	return c.JSON(&fiber.Map{
		"success": true,
		"message": "Successfully fetched product",
		"product": product,
	})
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
	// TODO: Add validation
	errorValidation := XValidator{validator: validate}

	// Instantiate new Product struct
	p := new(model.Product)

	// Parse body into product struct
	if err := c.BodyParser(p); err != nil {
		log.Println(err)
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": "Invalid request body",
		})
	}

	// Validate the product
	if errs := errorValidation.Validate(p); len(errs) > 0 {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": "Invalid request body",
			"errors":  errs,
		})
	}

	// Insert Product into database
	if err := database.DB.Create(p).Error; err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Error creating product",
		})
	}

	// Return Product in JSON format
	return c.JSON(&fiber.Map{
		"success": true,
		"message": "Product successfully created",
		"product": p,
	})
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

// UpdateProduct godoc
// @Summary Update a product
// @Description Update a product
// @Tags products
// @Accept */*
// @Produce json
// @Param id path int true "Product ID"
// @Param name body string true "Product Name"
// @Param description body string true "Product Description"
// @Param category body string true "Product Category"
// @Param amount body int true "Product Amount"
// @Success 200 {object} map[string]interface{}
// @Router /api/{id} [put]
func UpdateProduct(c *fiber.Ctx) error {
	// Extract product ID from the URL parameters
	id := c.Params("id")

	// Instantiate new Product struct
	updatedProduct := new(model.Product)

	// Parse request body into updated product struct
	if err := c.BodyParser(updatedProduct); err != nil {
		log.Println(err)
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": "Invalid request body",
		})
	}

	// Fetch existing product from the database
	existingProduct := new(model.Product)
	if err := database.DB.First(existingProduct, id).Error; err != nil {
		log.Println(err)
		return c.Status(404).JSON(&fiber.Map{
			"success": false,
			"message": "Product not found",
		})
	}

	// Update fields of the existing product with the new values
	existingProduct.Name = updatedProduct.Name
	existingProduct.Description = updatedProduct.Description
	existingProduct.Category = updatedProduct.Category
	existingProduct.Amount = updatedProduct.Amount

	// Save the updated product back to the database
	if err := database.DB.Save(existingProduct).Error; err != nil {
		log.Println(err)
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Error updating product",
		})
	}

	// Return the updated product in JSON format
	return c.JSON(&fiber.Map{
		"success": true,
		"message": "Product successfully updated",
		"product": existingProduct,
	})
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags health
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health [get]
func HealthCheck(c *fiber.Ctx) error {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	if err := c.JSON(res); err != nil {
		return err
	}

	return nil
}
