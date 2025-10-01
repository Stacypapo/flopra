package service

import (
	"flowershy/internal/models"
	"flowershy/internal/repository"
	"flowershy/pkg/jwt"
)

type User interface {
	Register(email, password string) (int64, error)
	Login(email, password string) (string, error)
	UserInfo(userId int64) (*models.User, error)
	UpdateUser(user *models.User) (int64, error)
	ListUsers(limit, offset int) ([]models.User, error)
	CountUsers() (int64, error)
}

type Order interface {
	CreateOrder(userId int64, items map[int64]int) (int64, error)
	GetOrderById(orderId int64) (*models.Order, error)
	GetOrdersByUserId(userId int64) ([]models.Order, error)
	GetOrdersByWarehouseId(warehouseId int64) ([]models.Order, error)
	GetAllOrders(limit, offset int) ([]models.Order, error)
	CountOrders() (int64, error)
	UpdateOrderStatus(orderId int64, status string) (int64, error)
	DeleteOrder(orderId int64) (int64, error)
}

type Product interface {
	CreateProduct(product *models.Product) (int64, error)
	GetProductById(productId int64) (*models.Product, error)
	GetProductByName(name string) ([]models.Product, error)
	GetProductBySKU(sku string) (*models.Product, error)
	SearchProducts(query string, tags []int64) ([]models.Product, error) // удобный поиск
	GetAllProducts(limit, offset int) ([]models.Product, error)
	CountProducts() (int64, error)
	UpdateProduct(product *models.Product) (int64, error)
	DeleteProduct(productId int64) (int64, error)
}

type Cart interface {
	AddToCart(userId, productId int64, quantity int) error
	RemoveFromCart(userId, productId int64) error
	GetCartItems(userId int64) ([]models.CartItem, error)
	ClearCart(userId int64) error
	CountCartItems(userId int64) (int64, error)
}

type Inventory interface {
	UpdateInventory(warehouseId, productId int64, quantity int) error
	GetInventory(productId int64) ([]models.Inventory, error)
	GetByWarehouse(warehouseId int64) ([]models.Inventory, error)
	ReserveStock(orderId int64, items map[int64]int) error
}

type Ai interface {
	GenerateBouquet(query string) (map[string]int, string, error)
	CreateGeneratedBouquet(flowers map[string]int) (int64, error)
}

type Tag interface {
	AddTag(name string) (int64, error)
	GetTagById(tagId int64) (*models.Tag, error)
	GetAllTags() ([]models.Tag, error)
}

type Bouquet interface {
	CreateBouquet(parentId int64, items map[int64]int) error
	GetBouquet(parentId int64) ([]models.BouquetItem, error)
	DeleteBouquet(parentId int64) error
}

type UserQuery interface {
	SaveQuery(userId, productId int64, query, model string) error
	GetUserQueries(userId int64) ([]models.UserQuery, error)
	GetAllQueries(limit, offset int) ([]models.UserQuery, error)
}

type Service struct {
	User
	Order
	Product
	Cart
	Inventory
	Ai
	Tag
	Bouquet
	UserQuery
	Warehouse
}

func NewService(repo *repository.Repository, jwt_manager *jwt.JWTManager) *Service {
	return &Service{
		User:      NewUserService(repo.User, jwt_manager),
		Order:     NewOrderService(repo.Order, repo.OrderItem, repo.Inventory),
		Product:   NewProductService(repo.Product, repo.ProductTag),
		Cart:      NewCartService(repo.CartItem),
		Inventory: NewInventoryService(repo.Inventory),
		Ai:        NewAiService(repo.Inventory, repo.Product),
		Tag:       NewTagService(repo.Tag),
		Bouquet:   NewBouquetService(repo.BouquetItem),
		UserQuery: NewUserQueryService(repo.UserQuery),
		Warehouse: NewWarehouseService(repo.Warehouse),
	}
}
