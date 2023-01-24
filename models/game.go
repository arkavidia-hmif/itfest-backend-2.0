package models

import "gorm.io/gorm"

type Game struct {
	gorm.Model
	UserID         uint `gorm:"not null"`
	CurrentClueId  int32
	RemainingTries int32
	ClueDone       []int32 `gorm:"type:integer[]"`
}
