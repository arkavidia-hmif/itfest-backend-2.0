package services

import (
	"itfest-backend-2.0/configs"
	"itfest-backend-2.0/models"
)

func NewGame(gid uint) (string, error) {
	// db := configs.DB.GetConnection()


	return "success", nil
}

func CreateGame(id uint) error {
	db := configs.DB.GetConnection()

	clueId, err := GetRandomClueId([]uint{})
	if err != nil {
		return err
	}

	newGame := models.Game{
		UserID:         id,
		CurrentClueId:  clueId,
		RemainingTries: 3,
	}
	if err := db.Create(&newGame).Error; err != nil {
		return err
	}

	return nil
}
