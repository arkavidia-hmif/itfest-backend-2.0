package services

import (
	"gorm.io/gorm"
	"itfest-backend-2.0/configs"
	"itfest-backend-2.0/models"
)

func GrantPoint(userId uint, point uint) (models.User, error) {
	db := configs.DB.GetConnection()
	user := models.User{}

	if err := db.First(&user, models.User{Model: gorm.Model{ID: userId}}).Error; err != nil {
		return user, err
	}

	err := db.Model(&user).Where("id = ?", user.ID).Update("point", user.Point+point).Error

	return user, err
}
