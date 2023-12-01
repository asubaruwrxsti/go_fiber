package router

import (
	"github.com/firebase007/go-rest-api-with-fiber/handler"
	"github.com/firebase007/go-rest-api-with-fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	// Public routes
	app.Get("/", handler.Home)
	app.Get("/auth", middleware.Auth())
	app.Get("/health", handler.HealthCheck)

	// Middleware
	api := app.Group("/api", middleware.AuthReq())

	api.Get("/", handler.GetAllProducts)
	api.Get("/:id", handler.GetSingleProduct)
	api.Post("/", handler.CreateProduct)
	api.Post("/:id", handler.UpdateProduct)
	api.Delete("/:id", handler.DeleteProduct)
}
