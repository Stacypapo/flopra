package repository

import (
	"database/sql"
	"flowershy/internal/models"
)

type User interface {
	Create(user *models.User) (int64, error)
	ReadById(id int64) (*models.User, error)
	ReadByEmail(email string) (*models.User, error)
	ReadAll(limit, offset int) ([]*models.User, error) // пагинация
	Count() (int64, error)
	Update(user *models.User) (int64, error)
	Delete(id int64) (int64, error)
}

type Order interface {
	Create(order *models.Order) (int64, error)
	ReadById(id int64) (*models.Order, error)
	ReadByUserId(userId int64, limit, offset int) ([]models.Order, error)
	ReadByWarehouseId(warehouseId int64, limit, offset int) ([]models.Order, error)
	ReadAll(limit, offset int) ([]models.Order, error)
	Count() (int64, error)
	CountByUserId(userId int64) (int64, error)
	Update(order *models.Order) (int64, error)
	Delete(id int64) (int64, error)
}

type Product interface {
	Create(product *models.Product) (int64, error)
	ReadById(id int64) (*models.Product, error)
	ReadByName(name string, limit, offset int) ([]models.Product, error)
	ReadBySKU(sku string) (*models.Product, error)
	ReadAll(limit, offset int) ([]models.Product, error)
	Count() (int64, error)
	Update(product *models.Product) (int64, error)
	Delete(id int64) (int64, error)
	Search(query string, tags []int64) ([]models.Product, error)
}

type OrderItem interface {
	Create(item *models.OrderItem) error
	ReadByOrderId(orderId int64) ([]*models.OrderItem, error)
	ReadByProductId(productId int64) ([]*models.OrderItem, error)
	Delete(orderId, productId int64) error
	DeleteByOrderId(orderId int64) error
}

type CartItem interface {
	Create(item *models.CartItem) error
	ReadByUserId(userId int64) ([]models.CartItem, error)
	UpdateQuantity(userId, productId int64, quantity int) error
	Delete(userId, productId int64) error
	ClearCart(userId int64) error
	CountByUserId(userId int64) (int64, error) // сколько товаров в корзине
}

type UserQuery interface {
	Create(query *models.UserQuery) error
	ReadByUserId(userId int64, limit, offset int) ([]*models.UserQuery, error)
	ReadByProductId(productId int64, limit, offset int) ([]*models.UserQuery, error)
	ReadAll(limit, offset int) ([]*models.UserQuery, error)
	Count() (int64, error)
	CountByUserId(userId int64) (int64, error)
}

type BouquetItem interface {
	Add(item *models.BouquetItem) error
	ReadByParentId(productId int64) ([]*models.BouquetItem, error)
	Delete(productIdParent, productIdChild int64) error
	DeleteByParentId(productIdParent int64) error
}

type Tag interface {
	Create(tag *models.Tag) (int64, error)
	ReadById(id int64) (*models.Tag, error)
	ReadAll(limit, offset int) ([]models.Tag, error)
	Count() (int64, error)
	Update(tag *models.Tag) (int64, error)
	Delete(id int64) (int64, error)
}

type ProductTag interface {
	Add(pt *models.ProductTag) error
	ReadByProductId(productId int64) ([]models.Tag, error)
	ReadByTagId(tagId int64) ([]models.Product, error)
	Delete(tagId, productId int64) error
	DeleteByProductId(productId int64) error
}

type Warehouse interface {
	Create(warehouse *models.Warehouse) (int64, error)
	ReadById(id int64) (*models.Warehouse, error)
	ReadAll(limit, offset int) ([]models.Warehouse, error)
	Count() (int64, error)
	Update(warehouse *models.Warehouse) (int64, error)
	Delete(id int64) (int64, error)
}

type Inventory interface {
	Add(item *models.Inventory) error
	ReadByProductId(productId int64) ([]models.Inventory, error)
	ReadByWarehouseId(warehouseId int64) ([]models.Inventory, error)
	UpdateQuantity(item *models.Inventory) error
	Delete(warehouseId, productId int64) error
	ReserveStock(orderId int64, items map[int64]int) error
}

type Repository struct {
	User
	Order
	Product
	OrderItem
	CartItem
	UserQuery
	BouquetItem
	Tag
	ProductTag
	Warehouse
	Inventory
}

// NewRepository создает новый экземпляр Repository.
func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		User:        NewUserPostgres(db),
		Order:       NewOrderPostgres(db),
		Product:     NewProductPostgres(db),
		OrderItem:   NewOrderItemPostgres(db),
		CartItem:    NewCartItemPostgres(db),
		UserQuery:   NewUserQueryPostgres(db),
		BouquetItem: NewBouqetItemPostgres(db),
		Tag:         NewTagPostgres(db),
		ProductTag:  NewProductTagPostgres(db),
		Warehouse:   NewWarehousePostgres(db),
		Inventory:   NewInventoryPostgres(db),
	}
}
