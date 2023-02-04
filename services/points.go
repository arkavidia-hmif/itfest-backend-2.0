package services

import (
	"gorm.io/gorm"
	"itfest-backend-2.0/configs"
	"itfest-backend-2.0/models"
)

func GrantPoint(fromId uint, toId uint, point uint) (models.User, error) {
	db := configs.DB.GetConnection()
	user := models.User{}
	log := models.Log{
		From:  fromId,
		To:    toId,
		Point: point,
	}

	if err := db.First(&user, models.User{Model: gorm.Model{ID: toId}}).Error; err != nil {
		return user, err
	}

	if err := db.Model(&user).Where("id = ?", user.ID).Update("point", user.Point+point).Error; err != nil {
		return user, err
	}

	if err := db.Create(&log).Error; err != nil {
		return user, err
	}

	return user, nil
}
