package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	"github.com/imran31415/example-project-proto-db/auth"
	"github.com/imran31415/example-project-proto-db/config"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	Db *sql.DB
	auth.UnimplementedAuthServiceServer
}

func connectToDB(cfg config.DatabaseConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping MySQL: %w", err)
	}

	log.Println("Successfully connected to the database.")
	return db, nil
}

func startGRPCServers(server *Server) {
	go func() {
		log.Println("Starting insecure gRPC server on port 50052...")
		insecureServer := grpc.NewServer()
		auth.RegisterAuthServiceServer(insecureServer, server)
		grpc_health_v1.RegisterHealthServer(insecureServer, health.NewServer())
		reflection.Register(insecureServer)

		listenAndServe(insecureServer, ":50052")
	}()

	// Block the main goroutine to prevent the program from exiting
	select {}
}

func listenAndServe(server *grpc.Server, port string) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server on port %s: %v", port, err)
	}
}
