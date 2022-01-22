package main

import (
	mid "organizations-service/pkg/middleware"
	"organizations-service/pkg/routes"
	"organizations-service/platform/database"
	"os"

	_ "organizations-service/docs"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/swagger/*", swagger.HandlerDefault)
	mid.FiberMiddleware(app)
	database.Init()
	routes.PublicRoutes(app)
	app.Listen(os.Getenv("PORT"))
}
