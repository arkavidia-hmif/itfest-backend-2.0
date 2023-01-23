package models

type User struct {
	AccountID uint `gorm:"not null"`
	Point     int32
}
