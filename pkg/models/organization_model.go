package models

import "github.com/kamva/mgm/v3"

type Organization struct {
	mgm.DefaultModel `bson:",inline"`
	OrganizationName string `db:"organization_name" json:"organization_name" validate:"required"`
	UserId           string `db:"user_id" json:"user_id" validate:"required"`
}

type CreateOrganization_Request struct {
	OrganizationName string `db:"organization_name" json:"organization_name" validate:"required"`
}
type EditOrganization_Request struct {
	OrganizationName string `db:"organization_name" json:"organization_name" validate:"required"`
	OrganizationId   string `db:"organization_id" json:"organization_id" validate:""`
}
type DeleteOrganization_Request struct {
	OrganizationId string `db:"organization_id" json:"organization_id" validate:""`
}

type CreateOrganization_Response struct {
	OrganizationId   string `db:"organization_id" json:"organization_id" validate:""`
	OrganizationName string `db:"name" json:"name" validate:"required"`
}
