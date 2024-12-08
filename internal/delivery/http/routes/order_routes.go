package routes

import (
	"gocommerce/internal/delivery/http/handler"
	"gocommerce/internal/usecase"

	"github.com/gorilla/mux"
)

func RegisterOrderRoutes(router *mux.Router, uc *usecase.OrderUsecase, midtransUsecase *usecase.MidtransUsecase, userUseCase *usecase.UserUseCase, productUsecase *usecase.ProductUseCase) {
	orderHandler := handler.NewOrderHandler(uc, midtransUsecase, userUseCase, productUsecase)
	router.HandleFunc("/orders", orderHandler.CreateOrder).Methods("POST")
	router.HandleFunc("/orders", orderHandler.GetOrders).Methods("GET")
}
