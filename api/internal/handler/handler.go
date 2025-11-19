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

	router.HandleFunc("POST /api/users/register", h.Register)
	router.HandleFunc("POST /api/users/login", h.Login)
	router.HandleFunc("GET /api/users/me", h.UserInfo)
	router.HandleFunc("PUT /api/users/me", h.UpdateUser)
	router.HandleFunc("DELETE /api/users/me", h.DeleteUser)

	router.HandleFunc("POST /api/products", h.CreateProduct)
	router.HandleFunc("GET /api/products/{id}", h.GetProductById)
	router.HandleFunc("GET /api/products", h.GetAllProducts)
	router.HandleFunc("PUT /api/products/{id}", h.UpdateProduct)
	router.HandleFunc("DELETE /api/products/{id}", h.DeleteProduct)
	router.HandleFunc("GET /api/products/search/{query}", h.SearchProducts)

	//router.HandleFunc("GET /api/cart", h.GetCartItems)
	//router.HandleFunc("POST /api/cart", h.AddToCart)
	//router.HandleFunc("DELETE /api/cart/{product_id}", h.RemoveFromCart)
	//router.HandleFunc("DELETE /api/cart", h.ClearCart)

	//router.HandleFunc("POST /api/orders", h.CreateOrder)
	//router.HandleFunc("GET /api/orders/{id}", h.GetOrderById)
	//router.HandleFunc("GET /api/orders", h.GetAllOrders)
	//router.HandleFunc("PUT /api/orders/{id}/status", h.UpdateOrderStatus)
	//router.HandleFunc("DELETE /api/orders/{id}", h.DeleteOrder)

	//router.HandleFunc("POST /api/tags", h.AddTag)
	//router.HandleFunc("GET /api/tags/{id}", h.GetTagById)
	//router.HandleFunc("GET /api/tags", h.GetAllTags)
	//router.HandleFunc("POST /api/tags/{tag_id}/products", h.GetProductsByTag)
	//router.HandleFunc("POST /api/tags/{tag_id}/products/{product_id}", h.ApplyTagToProduct)
	//router.HandleFunc("DELETE /api/tags/{tag_id}/products/{product_id}", h.RemoveTagFromProduct)
	//router.HandleFunc("GET /api/products/{product_id}/tags", h.GetTagsByProductId)

	//router.HandleFunc("POST /api/ai/generate_bouquet", h.GenerateBouquet)
	//router.HandleFunc("POST /api/ai/generate_bouquet_image", h.GenerateBouquetImage)

	//router.HandleFunc("GET /api/inventory", h.GetInventory)
	//router.HandleFunc("GET /api/inventory/warehouse/{warehouse_id}", h.GetByWarehouse)
	//router.HandleFunc("PUT /api/inventory", h.UpdateInventory)

	return router
}
