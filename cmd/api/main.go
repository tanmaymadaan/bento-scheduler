package main

import (
	"log"

	"github.com/tanmaymadaan/bento-scheduler/config"
	"github.com/tanmaymadaan/bento-scheduler/pkg/database"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Create Cassandra session
	cassandraSession, err := database.NewCassandraSession(cfg.CassandraHosts, cfg.CassandraKeyspace)
	if err != nil {
		log.Fatalf("Failed to create Cassandra session: %v", err)
	}
	defer cassandraSession.Close()

	// Test the connection
	if err := cassandraSession.TestConnection(); err != nil {
		log.Fatalf("Failed to connect to Cassandra: %v", err)
	}

	log.Println("Successfully connected to Cassandra")

	// Your API server code will go here
	log.Println("Application is ready to serve requests")
}
