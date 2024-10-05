package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	PublicHost             string
	Port                   string
	DBUser                 string
	DBPassword             string
	DBAddress              string
	DBName                 string
	JWTExpirationInSeconds int64
	JWTSecret              string
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		PublicHost:             getEnv("PUBLIC_HOST", "http://localhost"),
		Port:                   getEnv("PORT", "4200"),
		DBUser:                 getEnv("DB_USER", "root"),
		DBPassword:             getEnv("DB_PASSWORD", "admin"),
		DBAddress:              fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:                 getEnv("DB_NAME", "go_user_auth"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXP", 3600*24*7),
		JWTSecret:              getEnv("JWT_SECRET", "not-secret-secret-anymore?"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}
	return fallback
}