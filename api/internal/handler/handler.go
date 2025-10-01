package handler

import (
	"flowershy/internal/service"
	"net/http"
)

type Handler struct {
	services *service.Service
}

// NewHandler создает новый экземпляр Handler.
func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

// InitRoutes инициализирует маршруты HTTP для обработчика Handler и возвращает настроенный мультиплексор (*http.ServeMux).
func (h *Handler) InitRoutes() *http.ServeMux {
	router := http.NewServeMux()
	// Пример маршрута
	
	return router
}
