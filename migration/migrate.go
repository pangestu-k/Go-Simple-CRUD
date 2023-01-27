package main

import (
	"api-sekolah/initialize"
	"api-sekolah/models"
)

func init() {
	initialize.LoadDatabase()
}

func main() {
	//migrate tb user
	// initialize.DB.AutoMigrate(&models.Student{}, &models.Class{})
	initialize.DB.AutoMigrate(&models.Student{})
}
