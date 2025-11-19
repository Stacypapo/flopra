package handler

import (
	"encoding/json"
	"flowershy/internal/models"
	"fmt"
	"net/http"
	"strconv"
)

//router.HandleFunc("POST /api/users/register", h.Register)
//router.HandleFunc("POST /api/users/login", h.Login)
//router.HandleFunc("GET /api/users/me", h.UserInfo)
//router.HandleFunc("PUT /api/users/me", h.UpdateUser)
//router.HandleFunc("DELETE /api/users/me", h.DeleteUser)

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req models.User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	fmt.Println(req)
	if req.PhoneNumber == "" || req.Password == "" || req.Name == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	user, err := h.services.Register(req.Name, req.Email, req.PhoneNumber, req.Password)
	if err != nil {
		http.Error(w, "Error registering user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req struct {
		PhoneNumber string `json:"email"`
		Password    string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	tokens, err := h.services.Login(req.PhoneNumber, req.Password)
	if err != nil {
		http.Error(w, "Error logging in: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tokens)
}

func (h *Handler) UserInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	userID, ok := r.Context().Value("UserID").(string)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	user, err := h.services.UserInfo(int64(userIDInt))
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	safeUser := struct {
		ID          int64  `json:"id"`
		Email       string `json:"email"`
		Role        string `json:"role"`
		Name        string `json:"name"`
		PhoneNumber string `json:"phone"`
	}{
		ID:          user.UserId,
		Email:       user.Email,
		Role:        user.Role,
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(safeUser)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req models.User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	userID, ok := r.Context().Value("UserID").(string)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	req.UserId = int64(userIDInt)
	updatedUser, err := h.services.UpdateUser(&req)
	if err != nil {
		http.Error(w, "Error updating user: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedUser)
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	userID, ok := r.Context().Value("UserID").(string)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	deletedUser, err := h.services.DeleteUser(int64(userIDInt))
	if err != nil {
		http.Error(w, "Error deleting user: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(deletedUser)
}
