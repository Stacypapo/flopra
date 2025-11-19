package handler

import (
	"encoding/json"
	"flowershy/internal/models"
	"net/http"
	"strconv"
	"strings"
)

//router.HandleFunc("POST /api/products", h.CreateProduct)
//router.HandleFunc("GET /api/products/{id}", h.GetProductById)
//router.HandleFunc("GET /api/products", h.GetAllProducts)
//router.HandleFunc("PUT /api/products/{id}", h.UpdateProduct)
//router.HandleFunc("DELETE /api/products/{id}", h.DeleteProduct)
//router.HandleFunc("GET /api/products/search/{query}", h.SearchProducts)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req models.Product
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	productId, err := h.services.Product.CreateProduct(&req)
	if err != nil {
		http.Error(w, "Error creating product: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]int64{"product_id": productId})
}

func (h *Handler) GetProductById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	product, err := h.services.Product.GetProductById(id)
	if err != nil {
		http.Error(w, "Error fetching product: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (h *Handler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")
	limit, _ := strconv.Atoi(limitStr)
	offset, _ := strconv.Atoi(offsetStr)
	products, err := h.services.Product.GetAllProducts(limit, offset)
	if err != nil {
		http.Error(w, "Error fetching products: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req models.Product
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	req.ProductId = id
	updatedId, err := h.services.Product.UpdateProduct(&req)
	if err != nil {
		http.Error(w, "Error updating product: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]int64{"updated_product_id": updatedId})
}

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	deletedId, err := h.services.Product.DeleteProduct(id)
	if err != nil {
		http.Error(w, "Error deleting product: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]int64{"deleted_product_id": deletedId})
}
func (h *Handler) SearchProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	query := r.PathValue("query")
	tagsParam := r.URL.Query().Get("tags")
	var tags []int64
	if tagsParam != "" {
		tagStrs := splitAndTrim(tagsParam, ",")
		for _, ts := range tagStrs {
			tagId, err := strconv.ParseInt(ts, 10, 64)
			if err != nil {
				http.Error(w, "Invalid tag ID: "+ts, http.StatusBadRequest)
				return
			}
			tags = append(tags, tagId)
		}
	}
	products, err := h.services.Product.SearchProducts(query, tags)
	if err != nil {
		http.Error(w, "Error searching products: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
func splitAndTrim(s, sep string) []string {
	var result []string
	parts := strings.Split(s, sep)
	for _, p := range parts {
		trimmed := strings.TrimSpace(p)
		if trimmed != "" {
			result = append(result, trimmed)
		}
		return result
	}
	return result
}
