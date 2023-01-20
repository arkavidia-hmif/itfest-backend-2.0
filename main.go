package main

import (
	"itfest-backend-2.0/api"
	configs "itfest-backend-2.0/configs"
)

func main() {
	configs.ConnectDB()
	api.Run()
}
