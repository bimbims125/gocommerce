package main

import (
	"gocommerce/internal/delivery/http/routes"
	"gocommerce/internal/infra"
	"gocommerce/internal/repository"
	"gocommerce/internal/usecase"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	infra.InitDB()
	defer infra.DB.Close()

	userRepo := repository.NewUserRepository(infra.DB)
	UserUseCase := usecase.NewUserUseCase(userRepo)

	// Setup routes and start the server
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()
	routes.RegisterRoutes(subRouter, UserUseCase)

	log.Println("Server is running on port 3300")
	http.ListenAndServe(":3300", subRouter)
}
