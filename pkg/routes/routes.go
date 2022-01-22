package routes

import (
	"organizations-service/pkg/controllers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/api/v1")
	route.Post("/organization", controllers.PostCreateOrganization).Name("Create Organization")
	route.Put("/organization", controllers.PutEditOrganization).Name("Edit Organization")           // Este deberia ser /:id?
	route.Delete("/organization", controllers.DeleteRemoveOrganization).Name("Remove Organization") // Este deberia ser /:id?
	route.Get("/organization/:id", controllers.GetRetrieveOrganizationbyId).Name("Retrieve Organization By Id")
	route.Get("/organization", controllers.GetRetrieveOrganizationbyUserId).Name("Retrieve Organization By User Id")
	route.Post("/test", controllers.TestToken)
	route.Get("/", controllers.GetRoutes)
}
