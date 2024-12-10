package handler

import (
	"encoding/json"
	"gocommerce/internal/entity"
	"gocommerce/internal/usecase"
	"gocommerce/internal/utils"
	"net/http"
)

type CartHandler struct {
	usecase *usecase.CartUsecase
}

func NewCartHandler(usecase *usecase.CartUsecase) *CartHandler {
	return &CartHandler{usecase: usecase}
}

func (h *CartHandler) CreateCartHandler(w http.ResponseWriter, r *http.Request) {
	var cart entity.Cart
	if err := json.NewDecoder(r.Body).Decode(&cart); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.usecase.CreateCart(r.Context(), &cart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.JSONResponse(w, http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}
