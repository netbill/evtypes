package evtypes

import (
	"time"

	"github.com/google/uuid"
	"github.com/paulmach/orb"
)

const PlacesTopicV1 = "places.v1"

const PlaceCreatedEvent = "place.created"

type PlaceCreatedPayload struct {
	PlaceID        uuid.UUID `json:"place_id"`
	ClassID        uuid.UUID `json:"class_id"`
	OrganizationID uuid.UUID `json:"organization_id"`

	Status   string    `json:"status"`
	Verified bool      `json:"verified"`
	Point    orb.Point `json:"point"`
	Address  string    `json:"address"`
	Name     string    `json:"name"`

	Description *string `json:"description,omitempty"`
	IconKey     *string `json:"icon_key,omitempty"`
	BannerKey   *string `json:"banner_key,omitempty"`
	Website     *string `json:"website,omitempty"`
	Phone       *string `json:"phone,omitempty"`

	CreatedAt time.Time `json:"created_at"`
}

const PlaceUpdatedEvent = "place.updated"

type PlaceUpdatedPayload struct {
	PlaceID        uuid.UUID `json:"place_id"`
	ClassID        uuid.UUID `json:"class_id"`
	OrganizationID uuid.UUID `json:"organization_id"`

	Status   string `json:"status"`
	Verified bool   `json:"verified"`
	Address  string `json:"address"`
	Name     string `json:"name"`

	Description *string `json:"description,omitempty"`
	IconKey     *string `json:"icon_key,omitempty"`
	BannerKey   *string `json:"banner_key,omitempty"`
	Website     *string `json:"website,omitempty"`
	Phone       *string `json:"phone,omitempty"`

	Version   int32     `json:"version"`
	UpdatedAt time.Time `json:"updated_at"`
}

const PlaceDeletedEvent = "place.deleted"

type PlaceDeletedPayload struct {
	PlaceID   uuid.UUID `json:"place_id"`
	DeletedAt time.Time `json:"deleted_at"`
}
