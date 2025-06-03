package config

import (
	"log"
	"os"
)

type Config struct {
	Port        string
	JWTSecret   []byte
	MongoURI    string
	MongoDBName string
	MongoUser   string
	MongoPass   string
	RedisURI    string
}

func LoadConfig() *Config {
	return &Config{
		Port:        getEnv("API_PORT"),
		JWTSecret:   []byte(getEnv("JWT_SECRET")),
		MongoURI:    getEnv("MONGO_URI"),
		MongoDBName: getEnv("MONGO_DB_NAME"),
		MongoUser:   getEnv("MONGO_INITDB_ROOT_USERNAME"),
		MongoPass:   getEnv("MONGO_INITDB_ROOT_PASSWORD"),
		RedisURI:    getEnv("REDIS_URI"),
	}
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Printf("env %s not found", key)
	}

	return value
}
