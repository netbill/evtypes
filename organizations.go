package evtypes

import (
	"time"

	"github.com/google/uuid"
)

const OrganizationsTopicV1 = "organizations.v1"

const OrganizationCreatedEvent = "organization.created"

type OrganizationCreatedPayload struct {
	OrganizationID uuid.UUID `json:"organization_id"`
	Status         string    `json:"status"`
	Name           string    `json:"name"`
	IconKey        *string   `json:"icon_key,omitempty"`
	BannerKey      *string   `json:"banner_key,omitempty"`

	CreatedAt time.Time `json:"created_at"`
}

const OrganizationUpdatedEvent = "organization.updated"

type OrganizationUpdatedPayload struct {
	OrganizationID uuid.UUID `json:"organization_id"`
	Status         string    `json:"status"`
	Name           string    `json:"name"`
	IconKey        *string   `json:"icon_key,omitempty"`
	BannerKey      *string   `json:"banner_key,omitempty"`

	Version   int32     `json:"version"`
	UpdatedAt time.Time `json:"updated_at"`
}

const OrganizationDeletedEvent = "organization.deleted"

type OrganizationDeletedPayload struct {
	OrganizationID uuid.UUID `json:"organization_id"`
	DeletedAt      time.Time `json:"deleted_at"`
}
