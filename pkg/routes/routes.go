package routes

import (
	"organizations-service/pkg/controllers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/api/v1")
	route.Post("/organization", controllers.PostCreateOrganization)
	route.Put("/organization", controllers.PutEditOrganization)
	route.Get("/organization/:id", controllers.GetRetrieveOrganizationbyId)
	route.Get("/organization", controllers.GetRetrieveOrganizationbyUserId)
	route.Post("/test", controllers.TestToken)
	//route.Get("/", controllers.GetRoutes) //Vamos a usar esta ruta en raiz para que devuelva todas las rutas del microservicio, su metodo y su body
}
