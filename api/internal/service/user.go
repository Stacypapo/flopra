package service

import (
	"flowershy/internal/models"
	"flowershy/internal/repository"
	"flowershy/pkg/crypto"
	"flowershy/pkg/errors"
	"flowershy/pkg/jwt"
	"time"
)

type UserService struct {
	user_repo   repository.User
	jwt_manager *jwt.JWTManager
}

func NewUserService(user_repo repository.User, jwt_manager *jwt.JWTManager) *UserService {
	return &UserService{
		user_repo:   user_repo,
		jwt_manager: jwt_manager,
	}
}

func (s *UserService) Register(name string, email string, phone_number string, password string) (int64, error) {
	user := &models.User{
		Name:        name,
		Email:       email,
		PhoneNumber: phone_number,
		Password:    crypto.HashPassword(password),
		Role:        "user",
		CreatedAt:   time.Now(),
	}
	return s.user_repo.Create(user)
}

func (s *UserService) Login(email, password string) (map[string]string, error) {
	user, err := s.user_repo.ReadByEmail(email)
	if err != nil {
		return nil, err
	}
	if !crypto.CheckPasswordHash(password, user.Password) {
		return nil, errors.ErrInvalidPassword
	}
	tokens, err := s.jwt_manager.GenerateTokens(user.UserId, user.Role)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}

func (s *UserService) UserInfo(userId int64) (*models.User, error) {
	return s.user_repo.ReadById(userId)
}

func (s *UserService) UpdateUser(user *models.User) (int64, error) {
	return s.user_repo.Update(user)
}

func (s *UserService) CountUsers() (int64, error) {
	return s.user_repo.Count()
}

func (s *UserService) ListUsers(limit, offset int) ([]models.User, error) {
	users, err := s.user_repo.ReadAll(limit, offset)
	if err != nil {
		return nil, err
	}
	var res []models.User
	for _, u := range users {
		u.Password = ""
		res = append(res, *u)
	}
	return res, nil
}
