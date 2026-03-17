package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	PortGRPC int
}

func Load() (Config, error) {
	portStr := os.Getenv("GRPC_PORT")
	if portStr == "" {
		portStr = "50051"
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return Config{}, fmt.Errorf("invalid GRPC_PORT: %w", err)
	}
	return Config{PortGRPC: port}, nil
}
