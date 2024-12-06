package routes

import (
	"gocommerce/internal/delivery/http/handler"
	"gocommerce/internal/usecase"

	"github.com/gorilla/mux"
)

func RegisterCategoryRoutes(router *mux.Router, uc *usecase.CategoryUseCase) {
	categoryUsecase := usecase.NewCategoryUseCase(uc)

	categoryHandler := handler.NewCategoryHandler(categoryUsecase)

	router.HandleFunc("/categories", categoryHandler.CreateCategory).Methods("POST")
	router.HandleFunc("/categories", categoryHandler.GetCategories).Methods("GET")
}
