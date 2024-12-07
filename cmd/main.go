package main

import (
	"gocommerce/internal/config"
	"gocommerce/internal/delivery/http/routes"
	"gocommerce/internal/infra"
	"gocommerce/internal/repository"
	"gocommerce/internal/usecase"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.InitImageKitConfig()
	infra.InitDB()
	defer infra.DB.Close()

	// User dependencies
	userRepo := repository.NewUserRepository(infra.DB)
	UserUseCase := usecase.NewUserUseCase(userRepo)

	// Category dependencies
	categoryRepo := repository.NewCategoryRepository(infra.DB)
	CategoryUseCase := usecase.NewCategoryUseCase(categoryRepo)

	// Product dependencies
	productRepo := repository.NewProductRepository(infra.DB)
	ProductUseCase := usecase.NewProductUsecase(productRepo)
	// Setup routes and start the server
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()
	routes.RegisterUserRoutes(subRouter, UserUseCase)
	routes.RegisterCategoryRoutes(subRouter, CategoryUseCase)
	routes.RegisterProductRoutes(subRouter, ProductUseCase)

	log.Println("Server is running on port 3300")
	http.ListenAndServe(":3300", subRouter)
}
