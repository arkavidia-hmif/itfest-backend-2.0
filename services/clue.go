package services

func GetRandomClueId(clueDone []uint) (uint, error) {
	return 1, nil

	// db := configs.DB.GetConnection()

	// // get all clue rows
	// clue := db.First(&models.Clue{})
	// rows := int(clue.RowsAffected)

	// cluesNotDone := []int{}
	// for i := 1; i < rows; i++ {
	// 	if contains(clueDone, i) {
	// 		cluesNotDone = append(cluesNotDone, i)
	// 	}
	// }

	// randomIndex := rand.Intn(len(cluesNotDone))
	// pick := cluesNotDone[randomIndex]

	// return uint(pick), nil
}

func contains(s []uint, i int) bool {
	for _, v := range s {
		if int(v) == i {
			return true
		}
	}

	return false
}
