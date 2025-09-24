package repository

import (
	"database/sql"
	"flowershy/internal/models"
)

type User interface {
	Create(user *models.User) (int64, error)
	ReadById(id int64) (*models.User, error)
	Update(user *models.User) (int64, error)
	Delete(id int64) (int64, error)
}

type Order interface {
}

type Product interface {
	Create(user *models.Product) (int64, error)
	ReadById(id int64) (*models.Product, error)
	ReadByName(name string) (*models.Product, error)
	ReadBySKU(sku string) (*models.Product, error)
	Update(user *models.Product) (int64, error)
	Delete(id int64) (int64, error)
}

type OrderItem interface {
}

type CartItem interface {
}

type UserQuery interface {
}

type BouqetItem interface {
}

type Tag interface {
}

type ProductTag interface {
}

type Warehouse interface {
}

type Inventory interface {
}

type Repository struct {
	User
	Order
	Product
	OrderItem
	CartItem
	UserQuery
	BouqetItem
	Tag
	ProductTag
	Warehouse
	Inventory
}

// NewRepository создает новый экземпляр Repository.
func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		User:       NewUserPostgres(db),
		Order:      NewOrderPostgres(db),
		Product:    NewProductPostgres(db),
		OrderItem:  NewOrderItemPostgres(db),
		CartItem:   NewCartItemPostgres(db),
		UserQuery:  NewUserQueryPostgres(db),
		BouqetItem: NewBouqetItemPostgres(db),
		Tag:        NewTagPostgres(db),
		ProductTag: NewProductTagPostgres(db),
		Warehouse:  NewWarehousePostgres(db),
		Inventory:  NewInventoryPostgres(db),
	}
}
