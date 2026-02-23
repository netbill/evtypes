package evtypes

import (
	"time"

	"github.com/google/uuid"
)

const OrgRoleCreatedEvent = "role.created"

type OrgRoleCreatedPayload struct {
	RoleID         uuid.UUID `json:"role_id"`
	OrganizationID uuid.UUID `json:"organization_id"`

	Name        string    `json:"name"`
	Description string    `json:"description"`
	Color       string    `json:"color"`
	Version     int32     `json:"version"`
	CreatedAt   time.Time `json:"created_at"`

	Rank           int32     `json:"rank"`
	RanksRevision  int32     `json:"ranks_revision"`
	RanksUpdatedAt time.Time `json:"ranks_updated_at"`
}

const OrgRoleUpdatedEvent = "role.updated"

type OrgRoleUpdatedPayload struct {
	RoleID      uuid.UUID `json:"role_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Color       string    `json:"color"`
	Version     int32     `json:"version"`
	UpdatedAt   time.Time `json:"updated_at"`
}

const OrgRoleDeletedEvent = "role.deleted"

type OrgRoleDeletedPayload struct {
	RoleID    uuid.UUID `json:"role_id"`
	DeletedAt time.Time `json:"deleted_at"`

	// ranks changed as a result of deletion
	RanksRevision  int32     `json:"ranks_revision"`
	RanksUpdatedAt time.Time `json:"ranks_updated_at"`
}

const OrgRolesRanksReorderedEvent = "roles.ranks.reordered"

type RankMove struct {
	RoleID uuid.UUID `json:"role_id"`
	ToRank int32     `json:"to_rank"`
}

type OrgRolesRanksReorderedPayload struct {
	OrganizationID uuid.UUID  `json:"organization_id"`
	Moves          []RankMove `json:"moves"` // порядок важен!
	RanksRevision  int32      `json:"ranks_revision"`
	RanksUpdatedAt time.Time  `json:"ranks_updated_at"`
}

const OrgRolePermissionsUpdatedEvent = "role.permissions.updated"

type OrgRolePermissionsUpdatedPayload struct {
	RoleID  uuid.UUID   `json:"role_id"`
	Granted []uuid.UUID `json:"granted"` // что добавили
	Revoked []uuid.UUID `json:"revoked"` // что убрали

	PermissionsRevision  int32     `json:"permissions_revision"`
	PermissionsUpdatedAt time.Time `json:"permissions_updated_at"`
}

const OrgMemberRoleAddedEvent = "member_role.added"

type OrgMemberRoleAddedPayload struct {
	MemberID            uuid.UUID `json:"member_id"`
	RoleID              uuid.UUID `json:"role_id"`
	MemberRolesRevision int32     `json:"member_roles_revision"`
	AddedAt             time.Time `json:"added_at"`
}

const OrgMemberRoleRemovedEvent = "member_role.removed"

type OrgMemberRoleRemovedPayload struct {
	MemberID            uuid.UUID `json:"member_id"`
	RoleID              uuid.UUID `json:"role_id"`
	MemberRolesRevision int32     `json:"member_roles_revision"`
	RemovedAt           time.Time `json:"removed_at"`
}
