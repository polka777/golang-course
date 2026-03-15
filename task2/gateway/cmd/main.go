package main

import (
	"fmt"
	_ "gateway/docs"
	"gateway/internal/adapters"
	"gateway/internal/config"
	"gateway/internal/rest/handlers"
	"gateway/internal/usecase"
	"net/http"
	"os"
	"time"

	httpSwagger "github.com/swaggo/http-swagger"
)

//	@title			My GitHub Repository API
//	@version		1.0
//	@description	This is a server for get information about github repository

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080

func main() {

	cfg, err := config.Load()
	if err != nil {
		fmt.Println("failed load config")
		os.Exit(1)
	}

	client, err := adapters.NewGrpcClient(cfg.CollectorAddr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer client.Close()
	usecase := usecase.NewGatewayUsecase(client)
	handler := handlers.NewRepoHandler(usecase)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/repo", handler.GetRepositoryInfo)

	mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	server := http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.RESTPort),
		Handler:      mux,
		ReadTimeout:  6 * time.Second,
		WriteTimeout: 6 * time.Second,
	}
	fmt.Printf("Server started on port: %s\n", cfg.RESTPort)

	err = server.ListenAndServe()
	if err != nil {
		fmt.Println("Server listen error")
		os.Exit(1)
	}

}
