package main

import (
	"flowershy/internal/app"

	_ "flowershy/docs"
)

// @title FlowerShy API
// @version 1.0
// @description API для организации АИС магазина цветов
// @host localhost:8080
// @BasePath /
func main() {
	app.Run()
}
