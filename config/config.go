package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	USER_SERVICE    string
	API_GATEWAY     string
	REFRESH_TOKEN   string
	ACCESS_TOKEN    string
	PRODUCT_SERVICE string
	DEBT_SERVICE    string
}

func Load() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Print("No .env file found?")
	}

	config := Config{}
	config.USER_SERVICE = cast.ToString(Coalesce("USER_SERVICE", ":6006"))
	config.API_GATEWAY = cast.ToString(Coalesce("API_GATEWAY", ":1111"))
	config.REFRESH_TOKEN = cast.ToString(Coalesce("REFRESH_TOKEN", "secret"))
	config.ACCESS_TOKEN = cast.ToString(Coalesce("ACCESS_TOKEN", "secret"))
	config.PRODUCT_SERVICE = cast.ToString(Coalesce("PRODUCT_SERVICE", ":9091"))
	config.DEBT_SERVICE = cast.ToString(Coalesce("DEBT_SERVICE", ":8075"))

	return &config
}

func Coalesce(key string, defaultValue interface{}) interface{} {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
