package routes

import (
	"gocommerce/internal/delivery/http/handler"
	"gocommerce/internal/usecase"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, uc *usecase.UserUseCase) {
	userHandler := handler.NewUserHandler(uc)

	router.HandleFunc("/register", userHandler.Register).Methods("POST")
	router.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
}
