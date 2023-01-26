package models

import "gorm.io/gorm"

type Merchandise struct {
	gorm.Model
	Name  string
	Stock uint
	Point uint
}
