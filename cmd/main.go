package main

import (
	"flowershy/configs"
	"flowershy/internal/handler"
	"flowershy/internal/repository"
	"flowershy/internal/service"
	"log"
	"net/http"

	_ "flowershy/docs"
)

// @title FlowerShy API
// @version 1.0
// @description API для организации АИС магазина цветов
// @host localhost:8080
// @BasePath /
func main() {
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	db, err := repository.NewPostgresDB(config)
	if err != nil {
		log.Fatal(err)
	}
	if err := repository.Migrate(db); err != nil {
		log.Fatal(err)
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.InitRoutes()))
}
