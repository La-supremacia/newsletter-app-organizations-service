package controllers

import (
	"encoding/json"
	"fmt"
	"organizations-service/pkg/models"
	"organizations-service/pkg/services"
	"organizations-service/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"github.com/kamva/mgm/v3/operator"
	"go.mongodb.org/mongo-driver/bson"
)

func PostCreateOrganization(c *fiber.Ctx) error {
	u := new(models.CreateOrganization_Request)

	if err := c.BodyParser(u); err != nil {

		return err
	}
	userId := utils.GetClaim(c, "id")
	createdOrganization := services.New_Organization(userId, u.OrganizationName)

	err := mgm.Coll(createdOrganization).Create(createdOrganization)

	responseOrganization := services.New_CreateOrganization_Response(createdOrganization.ID.Hex(), createdOrganization.OrganizationName)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	fmt.Println("Successfully created a new Organization", responseOrganization)

	return c.Status(fiber.StatusOK).JSON(responseOrganization)
}

func PutEditOrganization(c *fiber.Ctx) error {
	u := new(models.EditOrganization_Request)

	if err := c.BodyParser(u); err != nil {
		return err
	}
	baseModel := &models.Organization{}
	coll := mgm.Coll(baseModel)
	_ = coll.FindByID(u.OrganizationId, baseModel)
	baseModel.OrganizationName = u.OrganizationName
	err := coll.Update(baseModel)

	responseOrganization := services.New_EditOrganization_Response(baseModel.ID.Hex(), baseModel.OrganizationName)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	fmt.Println("Successfully modified the Organization", responseOrganization)

	return c.Status(fiber.StatusOK).JSON(responseOrganization)
}
func DeleteRemoveOrganization(c *fiber.Ctx) error {
	u := new(models.DeleteOrganization_Request)

	if err := c.BodyParser(u); err != nil {
		return err
	}
	baseModel := &models.Organization{}
	coll := mgm.Coll(baseModel)
	_ = coll.FindByID(u.OrganizationId, baseModel)
	err := coll.Delete(baseModel)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	fmt.Println("Successfully removed the Organization")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"sucess":  true,
		"message": "The organization was successfully deleted",
	})
}

func GetRetrieveOrganizationbyId(c *fiber.Ctx) error {
	organization_id := c.Params("id")
	baseModel := &models.Organization{}
	coll := mgm.Coll(baseModel)
	err := coll.FindByID(organization_id, baseModel)

	responseOrganization := services.New_Organization(baseModel.UserId, baseModel.OrganizationName)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err)
	}
	fmt.Println("Successfully retrived the Organization", responseOrganization)

	return c.Status(fiber.StatusOK).JSON(baseModel)
}

func GetRetrieveOrganizationbyUserId(c *fiber.Ctx) error {
	user_id := utils.GetClaim(c, "id")
	baseModel := &models.Organization{}
	coll := mgm.Coll(baseModel)
	result := []models.Organization{}

	err := coll.SimpleFind(&result, bson.M{"userid": bson.M{operator.Eq: user_id}})

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err)
	}
	fmt.Println("Successfully retrived the Organization", result)

	return c.Status(fiber.StatusOK).JSON(result)
}

func TestToken(c *fiber.Ctx) error {
	name := utils.GetClaim(c, "id")
	return c.SendString("Welcome " + name)
}

func GetRoutes(c *fiber.Ctx) error {
	createOrganizationRoute := c.App().GetRoute("Create Organization")
	editOrganizationRoute := c.App().GetRoute("Edit Organization")
	deleteOrganizationRoute := c.App().GetRoute("Remove Organization")
	retrieveOrganizationByIdRoute := c.App().GetRoute("Retrieve Organization By Id")
	retrieveOrganizationByUserIdRoute := c.App().GetRoute("Retrieve Organization By User Id")

	createOrganization := services.New_CreateOrganization_Response("string", "string")
	editOrganization := services.New_EditOrganization_Response("string", "string")
	deleteOrganization := services.New_DeleteOrganization_Request("string")

	var ruts = [5]models.Route{
		services.NewRoute(createOrganizationRoute.Path, createOrganizationRoute.Method, createOrganizationRoute.Name, string(utils.First(json.Marshal(createOrganization)))),
		services.NewRoute(editOrganizationRoute.Path, editOrganizationRoute.Method, editOrganizationRoute.Name, string(utils.First(json.Marshal(editOrganization)))),
		services.NewRoute(deleteOrganizationRoute.Path, deleteOrganizationRoute.Method, deleteOrganizationRoute.Name, string(utils.First(json.Marshal(deleteOrganization)))),
		services.NewRoute(retrieveOrganizationByIdRoute.Path, retrieveOrganizationByIdRoute.Method, retrieveOrganizationByIdRoute.Name, "/:id=<organizationId>"),
		services.NewRoute(retrieveOrganizationByUserIdRoute.Path, retrieveOrganizationByUserIdRoute.Method, retrieveOrganizationByUserIdRoute.Name, ""),
	}
	return c.Status(fiber.StatusOK).JSON(ruts)
}
