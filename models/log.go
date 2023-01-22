package models

import "gorm.io/gorm"

type Log struct {
	gorm.Model
	From  string `gorm:"not null"`
	To    string `gorm:"not null"`
	Point int32  `gorm:"not null"`
}
