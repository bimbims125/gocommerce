package routes

import (
	"gocommerce/internal/delivery/http/handler"
	"gocommerce/internal/usecase"

	"github.com/gorilla/mux"
)

func RegisterCartRoutes(router *mux.Router, uc *usecase.CartUsecase) {
	cartHandler := handler.NewCartHandler(uc)

	router.HandleFunc("/carts", cartHandler.CreateCartHandler).Methods("POST")
	router.HandleFunc("/carts/{user_id}", cartHandler.GetCartByUserIDHandler).Methods("GET")
	// router.HandleFunc("/carts", cartHandler.GetCarts).Methods("GET")
}
