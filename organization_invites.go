package evtypes

import (
	"time"

	"github.com/google/uuid"
)

const OrgInviteCreatedEvent = "organization_invite.created"

type OrgInviteCreatedPayload struct {
	InviteID       uuid.UUID `json:"invite_id"`
	OrganizationID uuid.UUID `json:"organization_id"`
	AccountID      uuid.UUID `json:"account_id"`
	Status         string    `json:"status"`
	ExpiresAt      time.Time `json:"expires_at"`
	CreatedAt      time.Time `json:"created_at"`
}

const OrgInviteDeletedEvent = "organization_invite.deleted"

type OrgInviteDeletedPayload struct {
	InvitedID uuid.UUID `json:"invite_id"`
	DeletedAt time.Time `json:"deleted_at"`
}

const OrgInviteAcceptedEvent = "organization_invite.accepted"

type OrgInviteAcceptedPayload struct {
	InviteID   uuid.UUID `json:"invite_id"`
	AcceptedAt time.Time `json:"declined_at"`
}

const OrgInviteDeclinedEvent = "organization_invite.declined"

type OrgInviteDeclinedPayload struct {
	InviteID   uuid.UUID `json:"invite_id"`
	DeclinedAt time.Time `json:"declined_at"`
}
