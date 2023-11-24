package main

import (
	"github.com/gofiber/fiber/v2" // import the fiber package

	"log"

	"github.com/firebase007/go-rest-api-with-fiber/database"

	"github.com/firebase007/go-rest-api-with-fiber/router"

	_ "github.com/lib/pq"

	_ "go-rest-api-with-fiber/docs"

	"github.com/gofiber/swagger"
)

func main() { // entry point to our program

	// Connect to database
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New() // call the New() method - used to instantiate a new Fiber App

	// app.Use(middleware.Logger()) // use the Logger middleware

	router.SetupRoutes(app)
	// setup swagger
	app.Get("/swagger/*", swagger.HandlerDefault)     // default
	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://http://localhost:3000/doc.json",
		DeepLinking: false,
	}))

	app.Listen(":3000") // listen/Serve the new Fiber app on port 3000

}
