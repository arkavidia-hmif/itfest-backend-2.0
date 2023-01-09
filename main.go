package main

import (
	"github.com/joho/godotenv"
	"itfest-backend-2.0/api"
)

func main() {
	godotenv.Load()
	api.Run()
}
