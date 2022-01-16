package routes

import (
	"organizations-service/pkg/controllers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/api/v1")
	route.Post("/organization", controllers.PostCreateOrganization)
	route.Post("/test", controllers.TestToken)
	/*
		route.Delete("/organization/:id", controllers.PostCreateOrganization)
		route.Get("/organization/:id", controllers.PostCreateOrganization)
		route.Put("/organization", controllers.PostCreateOrganization) */
	//route.Get("/", controllers.GetRoutes) //Vamos a usar esta ruta en raiz para que devuelva todas las rutas del microservicio, su metodo y su body
	//route.Post("/sign-in", controllers.PostSignIn)
}
