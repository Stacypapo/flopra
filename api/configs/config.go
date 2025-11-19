package configs

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/xhit/go-str2duration/v2"
)

type Config struct {
	DBHost          string
	DBPort          string
	DBUsername      string
	DBPassword      string
	DBName          string
	DBSSLMode       string
	JWTSecretKey    string
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
	MLAPIURL        string
	MLAPIKey        string
}

// LoadConfig загружает конфигурацию из .env файла или переменных окружения
func LoadConfig() (Config, error) {
	_ = godotenv.Load()

	requiredVars := []string{"DB_HOST", "DB_PORT", "DB_USERNAME", "DB_PASSWORD", "DB_NAME", "JWT_SECRETKEY", "ML_API_URL", "ML_API_KEY"}
	for _, v := range requiredVars {
		if os.Getenv(v) == "" {
			log.Printf("Warning: Environment variable %s is not set\n", v)
		}
	}

	accessTokenTTL, err := str2duration.ParseDuration(getEnv("ACCESS_TOKEN_TTL", "15m"))
	if err != nil {
		return Config{}, err
	}
	refreshTokenTTL, err := str2duration.ParseDuration(getEnv("REFRESH_TOKEN_TTL", "168h"))
	if err != nil {
		return Config{}, err
	}
	return Config{
		DBHost:          getEnv("DB_HOST", "localhost"),
		DBPort:          getEnv("DB_PORT", "5432"),
		DBUsername:      getEnv("DB_USERNAME", "postgres"),
		DBPassword:      getEnv("DB_PASSWORD", "postgres"),
		DBName:          getEnv("DB_NAME", "postgres"),
		DBSSLMode:       getEnv("DB_SSLMODE", "disable"),
		JWTSecretKey:    getEnv("JWT_SECRETKEY", "mysecret"),
		AccessTokenTTL:  accessTokenTTL,
		RefreshTokenTTL: refreshTokenTTL,
		MLAPIURL:        getEnv("ML_API_URL", "http://localhost:5000"),
		MLAPIKey:        getEnv("ML_API_KEY", "secret_ml_api_key"),
	}, nil
}

// getEnv возвращает значение переменной окружения или значение по умолчанию
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
