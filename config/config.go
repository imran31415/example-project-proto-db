package config

import (
	"log"
	"os"
)

// DatabaseConfig represents the database configuration
type DatabaseConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DbName   string
}

// ServerConfig represents the server configuration
type ServerConfig struct {
	GRPCPort       string
	TLSCertPath    string
	TLSKeyPath     string
	TLSCaCertPath  string
	Environment    string
	GrpcGatewayURL string
}

// LoadConfig loads the database and server configurations
func LoadConfig() Config {
	return Config{
		Database: DatabaseConfig{
			User:     getEnv("DB_USER", "root"),
			Password: getEnv("DB_PASSWORD", "Password123!"),
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "3306"),
			DbName:   getEnv("DB_NAME", "example_project_proto_db"),
		},
		Server: ServerConfig{
			GRPCPort:       getEnv("GRPC_PORT", "50051"),
			TLSCertPath:    getEnv("TLS_CERT_PATH", "/etc/ssl/server.crt"),
			TLSKeyPath:     getEnv("TLS_KEY_PATH", "/etc/ssl/server.key"),
			TLSCaCertPath:  getEnv("TLS_CA_CERT_PATH", "/etc/ssl/ca.crt"),
			Environment:    getEnv("ENVIRONMENT", "development"),
			GrpcGatewayURL: getEnv("GRPC_GATEWAY_URL", "localhost:50052"),
		},
	}
}

// Config represents the overall configuration
type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
}

// getEnv fetches the value of an environment variable or returns a default value
func getEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Printf("Environment variable %s not set. Using default: %s", key, defaultValue)
		return defaultValue
	}
	return value
}
