package routes

import (
	"gocommerce/internal/delivery/http/handler"
	"gocommerce/internal/usecase"

	"github.com/gorilla/mux"
)

func RegisterOrderRoutes(router *mux.Router, uc *usecase.OrderUsecase) {
	orderHandler := handler.NewOrderHandler(uc)
	router.HandleFunc("/orders", orderHandler.CreateOrder).Methods("POST")
	router.HandleFunc("/orders", orderHandler.GetOrders).Methods("GET")
}
