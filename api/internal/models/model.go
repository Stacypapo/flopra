package models

import "time"

//Модель пользователей
type User struct {
	UserId      int64     `json:"user_id"`
	Email       string    `json:"email"`
	Name        string    `json:"name"`
	Password    string    `json:"password"`
	PhoneNumber string    `json:"phone"`
	Role        string    `json:"role"`
	CreatedAt   time.Time `json:"created_at"`
}

//Модель заказов
type Order struct {
	OrderId         int64     `json:"order_id"`
	UserId          int64     `json:"user_id"`
	Status          string    `json:"status"`
	TotalAmount     float64   `json:"total_amount"`
	CreatedAt       time.Time `json:"created_at"`
	ShippingAddress string    `json:"shipping_address"`
	WarehouseId     int64     `json:"warehouse_id"`
}

//Модель товаров
type Product struct {
	ProductId   int64     `json:"product_id"`
	SKU         string    `json:"sku"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	URL         string    `json:"url"`
}

//Модель товаров в заказе
type OrderItem struct {
	OrderId   int64 `json:"order_id"`
	ProductId int64 `json:"product_id"`
	Quantity  int   `json:"quantity"`
}

//Модель товаров в корзине
type CartItem struct {
	UserId    int64 `json:"user_id"`
	ProductId int64 `json:"product_id"`
	Quantity  int   `json:"quantity"`
}

//Модель запросов пользователей
type UserQuery struct {
	ProductId int64     `json:"product_id"`
	UserId    int64     `json:"user_id"`
	Query     string    `json:"query"`
	Model     string    `json:"model"`
	Rating    int       `json:"rating"`
	CreatedAt time.Time `json:"created_at"`
}

//Модель состава букетов
type BouquetItem struct {
	ProductIdParent int64 `json:"product_id_parent"`
	ProductIdChild  int64 `json:"product_id_child"`
	Quantity        int   `json:"quantity"`
}

//Модель тэгов
type Tag struct {
	TagId    int64  `json:"tag_id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	ColorHex string `json:"color_hex"`
}

//Модель тэгов к товару
type ProductTag struct {
	TagId     int64 `json:"tag_id"`
	ProductId int64 `json:"product_id"`
}

//Модель склада
type Warehouse struct {
	WarehouseId int64  `json:"warehouse_id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
}

//Модель наличия
type Inventory struct {
	WarehouseId int64 `json:"warehouse_id"`
	ProductId   int64 `json:"product_id"`
	Quantity    int   `json:"quantity"`
	Reserved    int   `json:"reserved"`
}
