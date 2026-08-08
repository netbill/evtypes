package evtypes

import (
	"time"

	"github.com/google/uuid"
)

type Profile struct {
	AccountID   uuid.UUID `json:"account_id"`
	Username    string    `json:"username"`
	Pseudonym   *string   `json:"pseudonym,omitempty"`
	Description *string   `json:"description,omitempty"`
	AvatarKey   *string   `json:"avatar_key,omitempty"`
	Version     int32     `json:"version"`

	UpdatedAt time.Time  `json:"updated_at"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

const ProfilesTopicV1 = "profiles.v1"

const ProfileCreatedEvent = "profile.created"

type ProfileCreatedPayload struct {
	Profile Profile `json:"profile"`
}

const ProfileUpdatedEvent = "profile.updated"

type ProfileUpdatedPayload struct {
	Profile Profile `json:"profile"`
}

const ProfileDeletedEvent = "profile.deleted"

type ProfileDeletedPayload struct {
	Profile Profile `json:"profile"`
}
