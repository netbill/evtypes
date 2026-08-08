package evtypes

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Pseudonym *string   `json:"pseudonym,omitempty"`
	Bio       *string   `json:"description,omitempty"`
	AvatarKey *string   `json:"avatar_key,omitempty"`
	Version   int32     `json:"version"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type UserEmail struct {
	UserID    uuid.UUID  `json:"User_id"`
	Email     string     `json:"email"`
	Verified  bool       `json:"verified"`
	Version   int32      `json:"version"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

const UsersTopicV1 = "Users.v1"

const UserCreatedEvent = "User.created"

type UserCreatedPayload struct {
	User      User      `json:"User"`
	UserEmail UserEmail `json:"User_email"`
}

const UserDeletedEvent = "User.deleted"

type UserDeletedPayload struct {
	User      User      `json:"User"`
	UserEmail UserEmail `json:"User_email"`
}
