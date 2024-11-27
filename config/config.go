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
	TASK_MANAGEMENT string
	REFRESH_KEY     string
	ACCESS_KEY      string
}

func Load() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Print("No .env file found?")
	}

	config := Config{}
	config.USER_SERVICE = cast.ToString(Coalesce("USER_SERVICE", ":6006"))
	config.API_GATEWAY = cast.ToString(Coalesce("API_GATEWAY", ":1111"))
	config.REFRESH_KEY = cast.ToString(Coalesce("REFRESH_KEY", "secret"))
	config.ACCESS_KEY = cast.ToString(Coalesce("ACCESS_KEY", "secret"))
	config.TASK_MANAGEMENT = cast.ToString(Coalesce("TASK_MANAGEMENT", ":8072"))

	return &config
}

func Coalesce(key string, defaultValue interface{}) interface{} {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
