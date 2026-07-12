package services

import (
	"barecms/internal/models"
	"barecms/internal/storage"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SiteUserService struct {
	storage *storage.Storage
}

func NewSiteUserService(storage *storage.Storage) *SiteUserService {
	return &SiteUserService{storage: storage}
}

// InviteUser invites a user to collaborate on a site
func (s *SiteUserService) InviteUser(siteID, inviterID, email, role string) (*models.InviteResponse, error) {
	// Validate role
	if role != models.RoleEditor && role != models.RoleViewer {
		return nil, errors.New("invalid role: must be 'editor' or 'viewer'")
	}

	// Check if inviter is site owner
	var site models.Site
	err := s.storage.DB.Where("id = ? AND owner_id = ?", siteID, inviterID).First(&site).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("site not found or insufficient permissions")
		}
		return nil, err
	}

	// Find user by email
	var user models.User
	err = s.storage.DB.Where("email = ? AND is_active = true", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	// Check if user is already a member
	var existingSiteUser models.SiteUser
	err = s.storage.DB.Where("site_id = ? AND user_id = ?", siteID, user.ID).First(&existingSiteUser).Error
	if err == nil {
		return nil, errors.New("user is already a member of this site")
	}

	// Create invitation
	siteUser := models.SiteUser{
		ID:        uuid.New().String(),
		SiteID:    siteID,
		UserID:    user.ID,
		Role:      role,
		InvitedBy: inviterID,
		InvitedAt: time.Now(),
		JoinedAt:  &time.Now(), // Auto-join for simplicity (no email confirmation)
	}

	err = s.storage.DB.Create(&siteUser).Error
	if err != nil {
		return nil, err
	}

	return &models.InviteResponse{
		ID:        siteUser.ID,
		Email:     user.Email,
		Role:      role,
		Status:    "joined",
		InvitedAt: siteUser.InvitedAt,
	}, nil
}

// GetSiteUsers returns all users for a site
func (s *SiteUserService) GetSiteUsers(siteID string) ([]models.SiteUserResponse, error) {
	var siteUsers []models.SiteUser
	err := s.storage.DB.Preload("User").Where("site_id = ?", siteID).Find(&siteUsers).Error
	if err != nil {
		return nil, err
	}

	// Also get the site owner
	var site models.Site
	err = s.storage.DB.Preload("Owner").Where("id = ?", siteID).First(&site).Error
	if err != nil {
		return nil, err
	}

	var response []models.SiteUserResponse

	// Add owner first
	ownerJoinedAt := site.CreatedAt
	response = append(response, models.SiteUserResponse{
		ID:       site.Owner.ID,
		User:     site.Owner,
		Role:     models.RoleOwner,
		JoinedAt: &ownerJoinedAt,
		Status:   "joined",
	})

	// Add collaborators
	for _, siteUser := range siteUsers {
		status := "pending"
		if siteUser.JoinedAt != nil {
			status = "joined"
		}

		response = append(response, models.SiteUserResponse{
			ID:       siteUser.ID,
			User:     siteUser.User,
			Role:     siteUser.Role,
			JoinedAt: siteUser.JoinedAt,
			Status:   status,
		})
	}

	return response, nil
}

// RemoveUser removes a user from a site (owner only)
func (s *SiteUserService) RemoveUser(siteID, ownerID, userToRemoveID string) error {
	// Verify owner permissions
	var site models.Site
	err := s.storage.DB.Where("id = ? AND owner_id = ?", siteID, ownerID).First(&site).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("site not found or insufficient permissions")
		}
		return err
	}

	// Cannot remove site owner
	if userToRemoveID == ownerID {
		return errors.New("cannot remove site owner")
	}

	// Remove from site_users
	result := s.storage.DB.Where("site_id = ? AND user_id = ?", siteID, userToRemoveID).Delete(&models.SiteUser{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("user not found in site")
	}

	return nil
}

// UpdateUserRole updates a user's role in a site (owner only)
func (s *SiteUserService) UpdateUserRole(siteID, ownerID, userID, newRole string) error {
	// Validate role
	if newRole != models.RoleEditor && newRole != models.RoleViewer {
		return errors.New("invalid role: must be 'editor' or 'viewer'")
	}

	// Verify owner permissions
	var site models.Site
	err := s.storage.DB.Where("id = ? AND owner_id = ?", siteID, ownerID).First(&site).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("site not found or insufficient permissions")
		}
		return err
	}

	// Cannot change owner role
	if userID == ownerID {
		return errors.New("cannot change owner role")
	}

	// Update role
	result := s.storage.DB.Model(&models.SiteUser{}).
		Where("site_id = ? AND user_id = ?", siteID, userID).
		Update("role", newRole)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("user not found in site")
	}

	return nil
}