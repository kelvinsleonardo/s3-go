package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"s3-microservice/common"
	"s3-microservice/interceptor"
	"s3-microservice/routes"
)

func main() {
	// Initialize logger
	err := common.InitLogger()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer common.CloseLogger()

	// Crie um roteador Mux
	router := mux.NewRouter()

	// Adicione o interceptor global
	router.Use(interceptor.GlobalInterceptor)

	// Configure as rotas
	routes.SetupPrometheusRoutes(router)
	routes.SetupUserRoutes(router)
	routes.SetupS3Routes(router)

	// Inicie o servidor HTTP
	//common.Logger.Info("Starting server on port 8080")
	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
