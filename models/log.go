package models

import "gorm.io/gorm"

type Log struct {
	gorm.Model
	From  uint  `gorm:"not null"`
	To    uint  `gorm:"not null"`
	Point int32 `gorm:"not null"`
}