package router

import (
	"github.com/firebase007/go-rest-api-with-fiber/config"
	"github.com/firebase007/go-rest-api-with-fiber/handler"
	"github.com/firebase007/go-rest-api-with-fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	// Public routes
	app.Get("/", handler.Home)

	// Middleware
	// api := app.Group("/api", middleware.AuthReq())
	api := app.Group("/api", middleware.JWTAuth(config.Config("JWT_SECRET")))

	api.Get("/", handler.GetAllProducts)
	api.Get("/:id", handler.GetSingleProduct)
	api.Post("/", handler.CreateProduct)
	api.Post("/:id", handler.UpdateProduct)
	api.Delete("/:id", handler.DeleteProduct)

	// Debug routes
	debug := app.Group("/debug")
	debug.Get("/ping", handler.Ping)
}
