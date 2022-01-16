package services

import "organizations-service/pkg/models"

func New_Organization(userId string, organizationName string) *models.Organization {
	return &models.Organization{
		UserId:           userId,
		OrganizationName: organizationName,
	}
}

func New_CreateOrganization_Response(organizationId string, organizationName string) *models.CreateOrganization_Response {
	return &models.CreateOrganization_Response{
		OrganizationId:   organizationId,
		OrganizationName: organizationName,
	}
}
