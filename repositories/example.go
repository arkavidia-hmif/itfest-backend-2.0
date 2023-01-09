package repositories

import model "itfest-backend-2.0/models"

func GetHelloWorld() model.ExampleModel {
	return model.ExampleModel{
		Message: "Hellow",
	}
}
