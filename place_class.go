package evtypes

import (
	"time"

	"github.com/google/uuid"
)

const PlaceClassesTopicV1 = "place.classes.v1"

const PlaceClassCreatedEvent = "place_class.created"

type PlaceClassCreatedPayload struct {
	PlaceClassID uuid.UUID  `json:"place_class_id"`
	ParentID     *uuid.UUID `json:"parent_id,omitempty"`
	Name         string     `json:"name"`
	Description  string     `json:"description"`
	IconKey      *string    `json:"icon_key,omitempty"`

	CreatedAt time.Time `json:"created_at"`
}

const PlaceClassUpdatedEvent = "place_class.updated"

type PlaceClassUpdatedPayload struct {
	PlaceClassID uuid.UUID  `json:"place_class_id"`
	ParentID     *uuid.UUID `json:"parent_id,omitempty"`
	Name         string     `json:"name"`
	Description  string     `json:"description"`
	IconKey      *string    `json:"icon_key,omitempty"`

	Version      int32      `json:"version"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeprecatedAt *time.Time `json:"deprecated_at,omitempty"`
}
