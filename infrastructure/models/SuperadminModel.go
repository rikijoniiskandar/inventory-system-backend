package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Superadmin struct {
	gorm.Model
	ID string `gorm:"primaryKey;type:char(36);before_create"`
	Username string `gorm:"unique"`
	Email string `gorm:"unique"`
	Password string
	Roles []Role `gorm:"many2many:superadmin_roles;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SuperadminSession struct {
	gorm.Model
	SuperadminID string
	Token        string `gorm:"unique"`
	ExpiresAt    time.Time
}

func (s *Superadmin) BeforeCreate(tx *gorm.DB) error {
    s.ID = uuid.New().String()
    return nil
}