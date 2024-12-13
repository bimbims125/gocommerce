package main

import (
	"gocommerce/internal/config"
	"gocommerce/internal/delivery/http/routes"
	"gocommerce/internal/infra"
	"gocommerce/internal/repository"
	"gocommerce/internal/usecase"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func main() {
	// Validator
	v := validator.New()

	// Init imagekit config
	config.InitImageKitConfig()

	// init database
	infra.InitDB()
	defer infra.DB.Close()

	userRepo := repository.NewUserRepository(infra.DB)
	UserUseCase := usecase.NewUserUseCase(userRepo)

	categoryRepo := repository.NewCategoryRepository(infra.DB)
	CategoryUseCase := usecase.NewCategoryUseCase(categoryRepo)

	productRepo := repository.NewProductRepository(infra.DB)
	ProductUseCase := usecase.NewProductUsecase(productRepo)

	orderRepo := repository.NewOrderRepository(infra.DB)
	orderUseCase := usecase.NewOrderUsecase(orderRepo)

	cartRepo := repository.NewCartRepository(infra.DB)
	cartUseCase := usecase.NewCartUsecase(*cartRepo)

	// Setup routes and start the server
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()
	routes.RegisterUserRoutes(subRouter, UserUseCase)
	routes.RegisterCategoryRoutes(subRouter, CategoryUseCase)
	routes.RegisterProductRoutes(subRouter, ProductUseCase, UserUseCase)
	routes.RegisterOrderRoutes(subRouter, orderUseCase, usecase.NewMidtransUsecase(v), UserUseCase, ProductUseCase)
	routes.RegisterMidtransRoutes(subRouter, usecase.NewMidtransUsecase(v), orderUseCase)
	routes.RegisterCartRoutes(subRouter, cartUseCase)

	log.Println("Server is running on port 3300")
	http.ListenAndServe(":3300", subRouter)
}
