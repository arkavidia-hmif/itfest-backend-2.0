package controllers

import (
	"itfest-backend-2.0/configs"
	"itfest-backend-2.0/models"
)

func createGame(id uint) error {
	db := configs.DB.GetConnection()

	newGame := models.Game{
		UserID:         id,
		CurrentClueId:  getRandomClueId(),
		RemainingTries: 3,
	}
	if err := db.Create(&newGame).Error; err != nil {
		return err
	}

	return nil
}
