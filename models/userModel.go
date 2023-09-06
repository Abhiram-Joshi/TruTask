package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"not null"`
	Password string
	IsAdmin  uint
	Active   uint
	Deleted  uint
	Roles    []*Role `gorm:"many2many:user_roles;"`
}
