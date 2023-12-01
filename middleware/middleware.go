package middleware

import (
	"log"
	"time"

	"github.com/firebase007/go-rest-api-with-fiber/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// Auth
func Auth() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		username := config.Config("AUTH_USERNAME")
		password := config.Config("AUTH_PASSWORD")
		log.Print(">> Inside AuthReq middleware")

		// Check if the username and password are correct
		if c.FormValue("username") == username && c.FormValue("password") == password {
			// Create a new token object, specifying signing method and the claims
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"username": username,
				"exp":      time.Now().Add(time.Hour * 72).Unix(),
			})

			// Sign and get the complete encoded token as a string using the secret
			tokenString, err := token.SignedString([]byte("your_secret"))
			if err != nil {
				return c.SendStatus(fiber.StatusInternalServerError)
			}
			log.Print("<< AuthReq middleware tokenString: ", tokenString)
			return c.JSON(fiber.Map{"token": tokenString})
		}
		log.Print(">> AuthReq middleware: username or password is incorrect")
		log.Printf(">> \nYour input username is: %s\nYour input password is %s", c.FormValue("username"), c.FormValue("password"))
		log.Printf(">> \nThe correct username is: %s\nThe correct password is %s", username, password)
		return c.SendStatus(fiber.StatusUnauthorized)
	}
}

// AuthReq middleware
func AuthReq() func(*fiber.Ctx) error {
	// check if the token is valid
	return func(c *fiber.Ctx) error {
		log.Print(">> Inside AuthReq middleware")
		token := c.Get("Authorization")
		if token == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		log.Print("<< AuthReq middleware: token: ", token)
		return c.Next()
	}
}
