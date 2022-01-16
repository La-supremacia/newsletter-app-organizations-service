package middleware

import (
	"organizations-service/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v3"
)

func FiberMiddleware(a *fiber.App) {
	a.Use(
		cors.New(),
		logger.New(),
	)

	secret := utils.GoDotEnvVariable("JWT_SECRET_KEY") // os.Getenv("JWT_SECRET_KEY")
	a.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(secret),
		ContextKey: "auth",
	}))

}
