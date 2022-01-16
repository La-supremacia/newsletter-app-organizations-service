package controllers

import (
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
