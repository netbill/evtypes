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

	Description *string `json:"description"`
	IconKey     *string `json:"icon_key"`
	BannerKey   *string `json:"banner_key"`
	Website     *string `json:"website"`
	Phone       *string `json:"phone"`

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

	Description *string `json:"description"`
	IconKey     *string `json:"icon_key"`
	BannerKey   *string `json:"banner_key"`
	Website     *string `json:"website"`
	Phone       *string `json:"phone"`

	Version   int32     `json:"version"`
	UpdatedAt time.Time `json:"updated_at"`
}

const PlaceDeletedEvent = "place.deleted"

type PlaceDeletedPayload struct {
	PlaceID   uuid.UUID `json:"place_id"`
	DeletedAt time.Time `json:"deleted_at"`
}
