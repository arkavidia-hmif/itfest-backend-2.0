package models

import "gorm.io/gorm"

type Merchandise struct {
	gorm.Model
	Name  string `json:"name"`
	Stock uint   `json:"stock"`
	Point uint   `json:"point"`
}
