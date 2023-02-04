package models

import "gorm.io/gorm"

type Clue struct {
	gorm.Model
	Usercode string `gorm:"not null"` // usercode for startups
	Text     string `gorm:"not null"`
}
