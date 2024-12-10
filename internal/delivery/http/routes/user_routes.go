package routes

import (
	"gocommerce/internal/delivery/http/handler"
	"gocommerce/internal/usecase"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(router *mux.Router, uc *usecase.UserUseCase) {
	userHandler := handler.NewUserHandler(uc)

	router.HandleFunc("/login", userHandler.Login).Methods("POST")
	router.HandleFunc("/register", userHandler.Register).Methods("POST")
	router.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", userHandler.GetUserByID).Methods("GET")
}
