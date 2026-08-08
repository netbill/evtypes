package evtypes

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID      uuid.UUID `json:"id"`
	Role    string    `json:"role"`
	Version int32     `json:"version"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type AccountEmail struct {
	AccountID uuid.UUID  `json:"account_id"`
	Email     string     `json:"email"`
	Verified  bool       `json:"verified"`
	Version   int32      `json:"version"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

const AccountsTopicV1 = "accounts.v1"

const AccountCreatedEvent = "account.created"

type AccountCreatedPayload struct {
	Account      Account      `json:"account"`
	AccountEmail AccountEmail `json:"account_email"`
}

const AccountDeletedEvent = "account.deleted"

type AccountDeletedPayload struct {
	Account      Account      `json:"account"`
	AccountEmail AccountEmail `json:"account_email"`
}
