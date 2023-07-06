package main

import (
	"ginTest/initialize"
	"ginTest/models"
)

func init() {
	initialize.LoadDotEnv()
	initialize.ConnectToDB()
}

func main() {
	err := initialize.DB.AutoMigrate(&models.User{}, &models.Review{})
	if err != nil {
		return
	}
}
