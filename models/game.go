package models

import "gorm.io/gorm"

type Game struct {
	gorm.Model
	UserID         uint `gorm:"not null"`
	CurrentClueId  uint
	RemainingTries uint
	ClueDone       []uint `gorm:"type:integer[]"`
}
