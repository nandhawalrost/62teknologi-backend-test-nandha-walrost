package main

import (
	"enamdua/initializers"
	"enamdua/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Business{})
}
