package models

import "gorm.io/gorm"

type Log struct {
	gorm.Model
	From  uint `gorm:"not null" json:"from"`
	To    uint `gorm:"not null" json:"to"`
	Point uint `gorm:"not null" json:"point"`
}
