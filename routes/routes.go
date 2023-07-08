package routes

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"s3-microservice/controllers"
)

func SetupPrometheusRoutes(router *mux.Router) {
	router.Handle("/metrics", promhttp.Handler())
}

func SetupUserRoutes(router *mux.Router) {
	userController := controllers.UserControllerInstance()
	router.HandleFunc("/users", userController.GetAllUsers).Methods("GET")
}

func SetupS3Routes(router *mux.Router) {
	s3Controller := controllers.S3ControllerInstance()
	router.HandleFunc("/upload", s3Controller.Upload).Methods("POST")
	router.HandleFunc("/download/{filename}", s3Controller.Download).Methods("GET")
}
