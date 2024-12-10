package routes

import (
	"gocommerce/internal/delivery/http/handler"
	"gocommerce/internal/infra"
	"gocommerce/internal/usecase"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(router *mux.Router, uc *usecase.ProductUseCase, userUsecase *usecase.UserUseCase) {
	productHandler := handler.NewProductHandler(uc)

	router.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	// router.HandleFunc("/products", productHandler.GetProducts).Methods("GET")
	router.HandleFunc("/products/{id}", productHandler.GetProductByID).Methods("GET")

	router.HandleFunc("/products", infra.WithJWTAuth(productHandler.GetProducts, *userUsecase)).Methods("GET")
}
