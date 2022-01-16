package main

import (
	mid "organizations-service/pkg/middleware"
	"os"

	"organizations-service/pkg/routes"
	"organizations-service/platform/database"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	mid.FiberMiddleware(app)
	database.Init()
	routes.PublicRoutes(app)
	app.Listen(os.Getenv("PORT"))
	//Commit para cerra issues
}
