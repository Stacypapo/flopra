package service

import (
	"context"
	"encoding/json"
	"flowershy/internal/models"
	"flowershy/internal/repository"
	"flowershy/pkg/errors"
	"fmt"
	"sync"
	"time"
)

type AiService struct {
	invRepo        repository.Inventory
	prodRepo       repository.Product
	user_queryRepo repository.UserQuery
	apiURL         string
	apiKey         string
}

func NewAiService(inv repository.Inventory, prod repository.Product, apiURL string, apiKey string) *AiService {
	return &AiService{
		invRepo:  inv,
		prodRepo: prod,
		apiURL:   apiURL,
		apiKey:   apiKey,
	}
}

type mlRequest struct {
	Query      string                 `json:"query"`
	Inventory  []models.Inventory     `json:"inventory"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}
//GenerateBouquetComposition(query string, warehouseId int64) (map[string]int, error)
//GenerateBouquetImage(flowers map[string]int) (string, error)
//CreateGeneratedBouquet(user_id int64, flowers map[string]int, url string) (int64, error)

func (s *AiService) GenerateBouquetComposition(query string, warehouseId int64) (map[string]int, error) {
	if query == "" {
		return nil, errors.ErrEmptyQuery
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var (
		inventory []models.Inventory
		errInv    error
	)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		inventory, errInv = s.invRepo.ReadByWarehouseId(warehouseId)
	}()

	reqCh := make(chan []byte, 1)
	errCh := make(chan error, 1)
	go func() {
		select {
		case <-ctx.Done():
			errCh <- ctx.Err()
		default:
			requestBody := mlRequest{
				Query:      query,
				Parameters: map[string]interface{}{"max_flowers": 8, "balance": true},
			}
			data, err := json.Marshal(requestBody)
			if err != nil {
				errCh <- err
				return
			}
			reqCh <- data
		}
	}()

	wg.Wait()
	if errInv != nil {
		return nil, fmt.Errorf("inventory error: %w", errInv)
	}

	select {
	case err := <-errCh:
		return nil, err
	case body := <-reqCh:
		var req mlRequest
		_ = json.Unmarshal(body, &req)
		req.Inventory = inventory

		finalBody, _ := json.Marshal(req)

		bouquet, err := s.CreateComposition(ctx, finalBody)
		return bouquet, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (s *AiService) GenerateBouquetImage(flowers map[string]int) (string, error) {
	if len(flowers) == 0 {
		return "", errors.New("EMPTY_BOUQUET", "empty bouquet")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	requestBody := mlRequest{
		Inventory:  []models.Inventory{},
		Parameters: map[string]interface{}{"generate_image": true},
	}
	requestBody.Parameters["bouquet"] = flowers

	data, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	imageURL, err := s.CreateImage(ctx, data)
	if err != nil {
		return "", err
	}
	return imageURL, nil
}

func (s *AiService) CreateGeneratedBouquet(user_id int64, flowers map[string]int, url string) (int64, error) {
	if len(flowers) == 0 {
		return 0, errors.New("EMPTY_BOUQUET", "empty bouquet")
	}

	var (
		mu          sync.Mutex
		price       float64
		description string = "Состав: "
		wg          sync.WaitGroup
		errCh       = make(chan error, len(flowers))
	)

	for name, qty := range flowers {
		wg.Add(1)
		go func(name string, qty int) {
			defer wg.Done()
			prods, err := s.prodRepo.ReadByName(name, 1, 0)
			if err != nil || len(prods) == 0 {
				errCh <- fmt.Errorf("flower %s not found", name)
				return
			}

			mu.Lock()
			price += prods[0].Price * float64(qty)
			description += fmt.Sprintf("%s x%d, ", name, qty)
			mu.Unlock()
		}(name, qty)
	}

	wg.Wait()
	close(errCh)
	for err := range errCh {
		if err != nil {
			return 0, err
		}
	}

	product := &models.Product{
		Name:        "Сгенерированный букет",
		Description: description[:len(description)-2],
		Price:       price,
		URL:         "https://cdn.flowershy.com/generated/default.jpg",
	}
	s.user_queryRepo.Create(&models.UserQuery{
		Query:     description,
		Model:     "composition-v1",
		CreatedAt: time.Now(),
		UserId:    user_id,
	})

	if url != "" {
		product.URL = url
	}

	return s.prodRepo.Create(product)
}

func (s *AiService) CreateComposition(ctx context.Context, body []byte) (map[string]int, error) {
	done := make(chan struct{})
	var result map[string]int
	var err error

	go func() {
		defer close(done)
		time.Sleep(400 * time.Millisecond)
		result = map[string]int{"Роза": 5, "Пион": 3, "Гербера": 2}
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-done:
		return result, err
	}
}

func (s *AiService) CreateImage(ctx context.Context, body []byte) (string, error) {
	done := make(chan struct{})
	var imageURL string
	var err error
	go func() {
		defer close(done)
		time.Sleep(700 * time.Millisecond)

		imageURL = "https://cdn.flowershy.com/generated/sample.jpg"
	}()
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-done:
		return imageURL, err
	}

}
