package services

import (
	"fmt"
	"math/rand"

	"github.com/lib/pq"
	"itfest-backend-2.0/configs"
	"itfest-backend-2.0/models"
)

func IsGameDone(game models.Game) bool {
	db := configs.DB.GetConnection()

	// get all clue rows
	clue := db.First(&models.Clue{})
	rows := int(clue.RowsAffected)

	if rows == len(game.CluesDone) {
		return true
	}

	return false
}

func GetRandomClueId(clueDone pq.Int32Array) (uint, error) {
	db := configs.DB.GetConnection()

	// get all clue rows
	clue := db.First(&models.Clue{})
	rows := int(clue.RowsAffected)

	cluesNotDone := pq.Int32Array{}
	for i := 1; i <= rows; i++ {
		if !contains(clueDone, i) {
			cluesNotDone = append(cluesNotDone, int32(i))
		}
	}

	randomIndex := rand.Intn(len(cluesNotDone))
	pick := cluesNotDone[randomIndex]

	return uint(pick), nil
}

func GetRandomClueIdByGame(game models.Game) (uint, error) {
	db := configs.DB.GetConnection()

	// get all clue rows
	var count int64
	db.Model(&models.Clue{}).Count(&count)
	rows := int(count)

	cluesDone := game.CluesDone
	fmt.Println(cluesDone)

	cluesNotDone := pq.Int32Array{}
	for i := 1; i <= rows; i++ {
		if !contains(cluesDone, i) {
			cluesNotDone = append(cluesNotDone, int32(i))
		}
	}

	// meaning that the game is done
	if len(cluesNotDone) == 0 {
		return 9999, nil
	}

	randomIndex := rand.Intn(len(cluesNotDone))
	pick := cluesNotDone[randomIndex]
	fmt.Println(pick)

	return uint(pick), nil
}

func contains(s pq.Int32Array, i int) bool {
	for _, v := range s {
		if int(v) == i {
			return true
		}
	}

	return false
}
