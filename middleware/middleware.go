package middleware

import (
	"github.com/firebase007/go-rest-api-with-fiber/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	jwtware "github.com/gofiber/jwt/v3"
)

// AuthReq middleware
func AuthReq() func(*fiber.Ctx) error {

	cfg := basicauth.Config{
		Users: map[string]string{
			config.Config("AUTH_USERNAME"): config.Config("AUTH_PASSWORD"),
		},
	}

	return basicauth.New(cfg)
}

// TODO: Add JWT middleware
func JWTAuth(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(secret),
	})
}
