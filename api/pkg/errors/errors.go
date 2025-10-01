package errors

import "fmt"

// AppError — кастомная ошибка с кодом
type AppError struct {
	Code    string // например, "USER_NOT_FOUND"
	Message string // человекочитаемое сообщение
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// Фабрики ошибок
func New(code, msg string) *AppError {
	return &AppError{Code: code, Message: msg}
}

// Примеры "преднастроенных" ошибок
var (
	ErrUserNotFound    = New("USER_NOT_FOUND", "user not found")
	ErrProductNotFound = New("PRODUCT_NOT_FOUND", "product not found")
	ErrOutOfStock      = New("OUT_OF_STOCK", "product is out of stock")
	ErrInvalidInput    = New("INVALID_INPUT", "invalid input provided")
	ErrInvalidPassword = New("INVALID_PASSWORD", "invalid password")
	ErrInvalidToken    = New("INVALID_TOKEN", "invalid token")
)
