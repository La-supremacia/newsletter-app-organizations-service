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

func New_EditOrganization_Response(organizationId string, organizationName string) *models.CreateOrganization_Response {
	return &models.CreateOrganization_Response{
		OrganizationId:   organizationId,
		OrganizationName: organizationName,
	}
}
func New_DeleteOrganization_Response(message string, success bool) *models.DeleteOrganization_Response {
	return &models.DeleteOrganization_Response{
		Message: message,
		Success: success,
	}
}
