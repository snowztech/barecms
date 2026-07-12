package models

import "time"

type Site struct {
	ID        string    `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name      string    `json:"name" gorm:"not null"`
	Slug      string    `json:"slug" gorm:"uniqueIndex;not null"`
	OwnerID   string    `json:"ownerId" gorm:"type:uuid;not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`

	// Relationships
	Owner     User       `json:"owner" gorm:"foreignKey:OwnerID"`
	SiteUsers []SiteUser `json:"siteUsers,omitempty" gorm:"foreignKey:SiteID"`
}

type CreateSiteRequest struct {
	Name    string `json:"name" validate:"required"`
	OwnerID string `json:"ownerId"` // Set from JWT context
}

type SiteData struct {
	ID   string                 `json:"id"`
	Name string                 `json:"name"`
	Slug string                 `json:"slug"`
	Data map[string]interface{} `json:"data"`
}
