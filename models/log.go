package models

import "gorm.io/gorm"

type Log struct {
	gorm.Model
	FromId uint `gorm:"not null" json:"from_id"`
	From   User `json:"from"`
	ToId   uint `gorm:"not null" json:"to_id"`
	To     User `json:"to"`
	Point  uint `gorm:"not null" json:"point"`
}
