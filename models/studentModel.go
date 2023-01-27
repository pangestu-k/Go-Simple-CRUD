package models

import "gorm.io/gorm"

// type Student struct {
// 	gorm.Model
// 	Name string `gorm:"size:50; not null"`
// 	// ClassID int
// 	// Class   Class
// 	NIM   string `gorm:"unique;size:50; not null"`
// 	Age   int    `gorm:"not null"`
// 	Grade string `gorm:"not null"`
// }

type Student struct {
	gorm.Model
	Name string
	// ClassID int
	// Class   Class
	NIM   string
	Age   int
	Grade string
}
