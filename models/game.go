package models

import "gorm.io/gorm"

type Game struct {
	gorm.Model
	UserId         string
	CurrentClueId  int32
	RemainingTries int32
	ClueDone       []int32
}
