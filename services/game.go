package services

import (
	"github.com/lib/pq"
	"itfest-backend-2.0/configs"
	"itfest-backend-2.0/models"
)

func NewGame(game models.Game) error {
	db := configs.DB.GetConnection()

	game.CluesDone = append(game.CluesDone, int32(game.CurrentClueId))
	clueId, err := GetRandomClueIdByGame(game)
	if err != nil {
		return err
	}

	if err := db.Find(&models.Game{}, game.ID).Updates(models.Game{
		CurrentClueId:  clueId,
		RemainingTries: 3,
		CluesDone:      game.CluesDone,
	}).Error; err != nil {
		return err
	}

	return nil
}

func CreateGame(id uint) error {
	db := configs.DB.GetConnection()

	clueId, err := GetRandomClueId(pq.Int32Array{})
	if err != nil {
		return err
	}

	newGame := models.Game{
		UserID:         id,
		CurrentClueId:  clueId,
		RemainingTries: 3,
		CluesDone:      pq.Int32Array{},
	}
	if err := db.Create(&newGame).Error; err != nil {
		return err
	}

	return nil
}
