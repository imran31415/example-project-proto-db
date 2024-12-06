package main

import (
	"log"

	"github.com/imran31415/example-project-proto-db/config"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to the database
	db, err := connectToDB(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize the gRPC server
	server := &Server{Db: db}

	// Start gRPC servers
	startGRPCServers(server)
}
