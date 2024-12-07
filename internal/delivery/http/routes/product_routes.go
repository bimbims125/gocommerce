package routes

import (
	"gocommerce/internal/delivery/http/handler"
	"gocommerce/internal/usecase"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(router *mux.Router, uc *usecase.ProductUseCase) {
	productHandler := handler.NewProductHandler(uc)

	router.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
}
