package app

import (
	"flowershy/configs"
	"flowershy/internal/handler"
	"flowershy/internal/repository"
	"flowershy/internal/service"
	"flowershy/pkg/jwt"
	"log"
	"net/http"

	_ "flowershy/docs"
)

// @title Payment System API
// @version 1.0
// @description API для управления транзакциями и кошельками
// @host localhost:8080
// @BasePath /
func Run() {
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
	jwt_manager := jwt.NewJWTManager(config.JWTSecretKey, config.AccessTokenTTL, config.RefreshTokenTTL)
	repos := repository.NewRepository(db)
	services := service.NewService(repos, jwt_manager, config.MLAPIURL, config.MLAPIKey)
	handlers := handler.NewHandler(services)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.InitRoutes()))
}
