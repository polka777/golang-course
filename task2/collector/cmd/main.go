package main

import (
	"collector/internal/adapters/github"
	"collector/internal/adapters/grpsserver"
	"collector/internal/config"
	"collector/internal/usecase"
	"fmt"
	"net"
	"os"

	pb "collector/proto"

	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Println("config!!!")
		os.Exit(1)
	}

	client := github.NewClient()
	getRepInfoUseCase := usecase.NewGetRepositoryInfoUseCase(client)
	server := grpc.NewServer()
	handler := grpsserver.NewServer(getRepInfoUseCase)
	pb.RegisterCollectorServiceServer(server, handler)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.PortGRPC))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)

	}
	fmt.Printf("Collector Server started serve on port: %d", cfg.PortGRPC)
	err = server.Serve(lis)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)

	}
}
