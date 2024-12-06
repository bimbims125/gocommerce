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

	// User dependencies
	userRepo := repository.NewUserRepository(infra.DB)
	UserUseCase := usecase.NewUserUseCase(userRepo)

	// Category dependencies
	categoryRepo := repository.NewCategoryRepository(infra.DB)
	CategoryUseCase := usecase.NewCategoryUseCase(categoryRepo)

	// Setup routes and start the server
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()
	routes.RegisterUserRoutes(subRouter, UserUseCase)
	routes.RegisterCategoryRoutes(subRouter, CategoryUseCase)

	log.Println("Server is running on port 3300")
	http.ListenAndServe(":3300", subRouter)
}
