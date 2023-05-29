package models

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name string
	Permissions []Permission `gorm:"many2many:role_permissions;"`
	Superadmins []Superadmin `gorm:"many2many:superadmin_roles;"`
}

type Permission struct {
	gorm.Model
	Name string
	Roles []Role `gorm:"many2many:role_permissions;"`
}