package service

import (
	//"flowershy/internal/models"
	"flowershy/internal/models"
	"flowershy/internal/repository"
)

type User interface {
	Register(username, email, password string) (int64, error)
	Login(email, password string) (string, error) // return token
	UserInfo(userId int64) (*models.User, error)
	UpdateUser(user *models.User) (int64, error)
}

type Order interface {
	CreateOrder(userId int64, items map[int64]int) (int64, error) // items map[product_id]count
	GetOrderById(orderId int64) (*models.Order, error)
	GetOrdersByUserId(userId int64) ([]models.Order, error)
	UpdateOrderStatus(orderId int64, status string) (int64, error)
	DeleteOrder(orderId int64) (int64, error)
}

type Product interface {
	CreateProduct(product *models.Product) (int64, error)
	GetProductById(productId int64) (*models.Product, error)
	GetProductByName(name string) (*models.Product, error)
	GetProductBySKU(sku string) (*models.Product, error)
	UpdateProduct(product *models.Product) (int64, error)
	DeleteProduct(productId int64) (int64, error)
}

type Cart interface {
	AddToCart(userId, productId int64, quantity int) (int64, error)
	RemoveFromCart(userId, productId int64) (int64, error)
	GetCartItems(userId int64) ([]models.CartItem, error)
	ClearCart(userId int64) (int64, error)
}

type Inventory interface {
	UpdateInventory(productId int64, quantity int) (int64, error)
	GetInventory(productId int64) (int, error)
}

type Ai interface {
	GenerateBouquet(query string) (map[string]int, string, error) // map[flower_name]count, url, error
	CreateGeneratedBouqet(flowers map[string]int) (int64, error)  // id, error
}

type Tag interface {
	AddTag(name string) (int64, error)
	GetTagById(tagId int64) (*models.Tag, error)
}

type Service struct {
	User
	Order
	Product
	Inventory
	Ai
}

// NewService создает новый экземпляр Service.
func NewService(repo *repository.Repository) *Service {
	return &Service{
		User:      NewUserService(repo.User),
		Order:     NewOrderService(repo.Order),
		Product:   NewProductService(repo.Product),
		Inventory: NewInventoryService(repo.Inventory),
		Ai:        NewAiService(),
		Tag:       NewTagService(repo.Tag),
	}
}
