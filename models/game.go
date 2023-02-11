package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	UserID         uint `gorm:"not null"`
	CurrentClueId  uint
	RemainingTries uint
	CluesDone      pq.Int32Array `gorm:"type:integer[]"`
}
