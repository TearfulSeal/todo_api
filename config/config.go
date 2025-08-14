package config

import(
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct{
	DBPath string
	ServerPort string
	JWTSecret string
}

func LoadConfig() *Config{
	err := godotenv.Load()
	if err != nil{
		log.Fatalf("failed to load .env file %v",err)
	}
	cfg := &Config{
		DBPath: getEnv("DB_PATH","todo.db"),
		ServerPort: getEnv("SERVER_PORT","8080"),
		JWTSecret: getEnv("JWT_SECRET","secret"),
	}
	return cfg

}

func getEnv(key, defaultValue string) string{
	if value, exists := os.LookupEnv(key); exists{
		return value
	}
	return defaultValue
}
