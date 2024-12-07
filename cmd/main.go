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

	userRepo := repository.NewUserRepository(infra.DB)
	categoryRepo := repository.NewCategoryRepository(infra.DB)
	productRepo := repository.NewProductRepository(infra.DB)
	orderRepo := repository.NewOrderRepository(infra.DB)

	UserUseCase := usecase.NewUserUseCase(userRepo)
	CategoryUseCase := usecase.NewCategoryUseCase(categoryRepo)
	ProductUseCase := usecase.NewProductUsecase(productRepo)
	orderUseCase := usecase.NewOrderUsecase(orderRepo)
	// Setup routes and start the server
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()
	routes.RegisterUserRoutes(subRouter, UserUseCase)
	routes.RegisterCategoryRoutes(subRouter, CategoryUseCase)
	routes.RegisterProductRoutes(subRouter, ProductUseCase)
	routes.RegisterOrderRoutes(subRouter, orderUseCase)

	log.Println("Server is running on port 3300")
	http.ListenAndServe(":3300", subRouter)
}
