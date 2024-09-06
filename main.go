package main

import (
	"week7/config"
	"week7/models"
	"week7/routes"
)

func main() {
	config.InitDB()
	models.MigrateDB()

	e := routes.InitRoutes()
	e.Start(":8080")
}
