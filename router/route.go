package router

import (
	"github.com/firebase007/go-rest-api-with-fiber/handler"

	"github.com/firebase007/go-rest-api-with-fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {

	// Middleware
	api := app.Group("/api", middleware.AuthReq())

	// routes

	// @ GET /api
	// @host localhost:3000
	// @desc Get all products
	// @version 1
	api.Get("/", handler.GetAllProducts)

	// @ GET /api/:id
	// @host localhost:3000
	// @desc Get single product
	// @version 1
	api.Get("/:id", handler.GetSingleProduct)

	// @ POST /api
	// @host localhost:3000
	// @desc Create a new product
	// @version 1
	api.Post("/", handler.CreateProduct)

	// @ DELETE /api/:id
	// @host localhost:3000
	// @desc Delete a product
	// @version 1
	api.Delete("/:id", handler.DeleteProduct)
}
