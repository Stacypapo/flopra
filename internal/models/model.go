package models

import "time"

//Модель пользователей
type User struct {
	UserId      int64     `json:"user_id"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	PhoneNumber string    `json:"phone"`
	Role        string    `json:"role"`
	CreatedAt   time.Time `json:"created_at"`
}

//Модель заказов
type Order struct {
	OrderID         int64     `json:"order_id"`
	UserID          int64     `json:"user_id"`
	Status          string    `json:"status"`
	TotalAmount     float64   `json:"total_amount"`
	CreatedAt       time.Time `json:"created_at"`
	ShippingAddress string    `json:"shipping_address"`
}

//Модель товаров
type Product struct {
	ProductID   int64     `json:"product_id"`
	SKU         string    `json:"sku"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	URL         string    `json:"url"`
}

//Модель товаров в заказе
type OrderItem struct {
	OrderID   int64 `json:"order_id"`
	ProductID int64 `json:"product_id"`
	Quantity  int   `json:"quantity"`
}

//Модель товаров в корзине
type CartItem struct {
	UserID    int64 `json:"user_id"`
	ProductID int64 `json:"product_id"`
	Quantity  int   `json:"quantity"`
}

//Модель запросов пользователей
type UserQuery struct {
	ProductID int64     `json:"product_id"`
	UserID    int64     `json:"user_id"`
	Query     string    `json:"query"`
	Model     string    `json:"model"`
	Rating    int       `json:"rating"`
	CreatedAt time.Time `json:"created_at"`
}

//Модель состава букетов
type BouqetItem struct {
	ProductIDParent int64 `json:"product_id_parent"`
	ProductIDChild  int64 `json:"product_id_child"`
	Quantity        int   `json:"quantity"`
}

//Модель тэгов
type Tag struct {
	TagID    int64  `json:"tag_id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	ColorHex string `json:"color_hex"`
}

//Модель тэгов к товару
type ProductTag struct {
	TagID     int64 `json:"tag_id"`
	ProductID int64 `json:"product_id"`
}

//Модель склада
type Warehouse struct {
	WarehouseID int64  `json:"warehouse_id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
}

//Модель наличия
type Inventory struct {
	WarehouseID int64 `json:"warehouse_id"`
	ProductID   int64 `json:"product_id"`
	Quantity    int   `json:"quantity"`
}
