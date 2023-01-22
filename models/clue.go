package models

import "gorm.io/gorm"

type Clue struct {
	gorm.Model
	Text string
}
