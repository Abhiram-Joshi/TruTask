package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	RoleID      uint
	Name        string
	Description string
	Done        uint
}
