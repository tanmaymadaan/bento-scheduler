package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	CassandraHosts    []string
	CassandraKeyspace string
	ServerPort        string
}

func LoadConfig() (*Config, error) {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	return &Config{
		CassandraHosts:    strings.Split(getEnv("CASSANDRA_HOSTS", "localhost:9042"), ","),
		CassandraKeyspace: getEnv("CASSANDRA_KEYSPACE", "bento_scheduler"),
		ServerPort:        getEnv("SERVER_PORT", "8080"),
	}, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
