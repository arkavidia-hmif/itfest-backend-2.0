package models

import "gorm.io/gorm"

type Clue struct {
	gorm.Model
	userCode string `gorm:"not null"`
	Text     string `gorm:"not null"`
}
