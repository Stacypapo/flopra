package repository

import (
	"flowershy/internal/models"
)

type User interface {
	Create(user *models.User) (int64, error)
	ReadById(id int64) (*models.User, error)
	Update(user *models.User) (int64, error)
	Delete(id int64) (int64, error)
}

// Модель заказов
type Order struct {
}

// Модель товаров
type Product struct {
}

// Модель товаров в заказе
type OrderItem struct {
}

// Модель товаров в корзине
type CartItem struct {
}

// Модель запросов пользователей
type UserQuery struct {
}

// Модель состава букетов
type BouqetItem struct {
}

// Модель тэгов
type Tag struct {
}

// Модель тэгов к товару
type ProductTag struct {
}

// Модель склада
type Warehouse struct {
}

// Модель наличия
type Inventory struct {
}
