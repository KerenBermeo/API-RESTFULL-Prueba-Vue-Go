package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model

	Title  string
	Done   bool `gorm:"default:false"`
	UserID uint
	Tasks  []Task
}
