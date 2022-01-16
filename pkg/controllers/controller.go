package controllers

import (
	"fmt"
	"organizations-service/pkg/models"
	"organizations-service/pkg/services"
	"organizations-service/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
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

func TestToken(c *fiber.Ctx) error {
	name := utils.GetClaim(c, "id")
	return c.SendString("Welcome " + name)
}
