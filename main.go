package main

import (
	"github.com/gofiber/fiber/v2" // import the fiber package

	"log"

	"github.com/firebase007/go-rest-api-with-fiber/database"
	"github.com/firebase007/go-rest-api-with-fiber/router"

	_ "github.com/lib/pq"

	_ "github.com/firebase007/go-rest-api-with-fiber/docs"

	"github.com/gofiber/swagger"
)

// @title Go Fiber REST API with Swagger Example
// @description Go Fiber REST API with Swagger Example
// @version 1
// @host localhost:3000
// @BasePath /
// @schemes http
func main() {

	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New() // call the New() method - used to instantiate a new Fiber App

	// app.Use(middleware.Logger())
	router.SetupRoutes(app)

	app.Get("/swagger/*", swagger.HandlerDefault)     // default
	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://localhost:3000/swagger/doc.json",
		DeepLinking: false,
	}))

	app.Listen(":3000")

}
