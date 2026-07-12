package models

import "time"

// SiteUser represents the many-to-many relationship between users and sites
type SiteUser struct {
	ID        string     `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	SiteID    string     `json:"siteId" gorm:"type:uuid;not null"`
	UserID    string     `json:"userId" gorm:"type:uuid;not null"`
	Role      string     `json:"role" gorm:"not null;default:'editor'"` // owner, editor, viewer
	InvitedBy string     `json:"invitedBy" gorm:"type:uuid"`
	InvitedAt time.Time  `json:"invitedAt" gorm:"autoCreateTime"`
	JoinedAt  *time.Time `json:"joinedAt"` // nil if invitation pending
	CreatedAt time.Time  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"autoUpdateTime"`

	// Relationships
	Site      Site `json:"site" gorm:"foreignKey:SiteID"`
	User      User `json:"user" gorm:"foreignKey:UserID"`
	Inviter   User `json:"inviter" gorm:"foreignKey:InvitedBy"`
}

// Ensure unique site-user combination
func (SiteUser) TableName() string {
	return "site_users"
}

// Role constants
const (
	RoleOwner  = "owner"
	RoleEditor = "editor"
	RoleViewer = "viewer"
)

// Role permissions
func (r string) CanEdit() bool {
	return r == RoleOwner || r == RoleEditor
}

func (r string) CanInvite() bool {
	return r == RoleOwner
}

func (r string) CanManageUsers() bool {
	return r == RoleOwner
}

// Invitation request models
type InviteUserRequest struct {
	Email string `json:"email" validate:"required,email"`
	Role  string `json:"role" validate:"required,oneof=editor viewer"`
}

type InviteResponse struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Status   string `json:"status"` // "pending" or "joined"
	InvitedAt time.Time `json:"invitedAt"`
}

type SiteUserResponse struct {
	ID       string     `json:"id"`
	User     User       `json:"user"`
	Role     string     `json:"role"`
	JoinedAt *time.Time `json:"joinedAt"`
	Status   string     `json:"status"` // "joined" or "pending"
}