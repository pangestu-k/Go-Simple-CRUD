package models

import "gorm.io/gorm"

type Class struct {
	gorm.Model
	Name string `gorm:"size:30;not null"`
}
