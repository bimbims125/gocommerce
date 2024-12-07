package handler

import (
	"encoding/json"
	"fmt"
	"gocommerce/internal/entity"
	"gocommerce/internal/usecase"
	"gocommerce/internal/utils"
	"net/http"
)

type OrderHandler struct {
	usecase *usecase.OrderUsecase
}

func NewOrderHandler(usecase *usecase.OrderUsecase) *OrderHandler {
	return &OrderHandler{usecase: usecase}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order entity.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := h.usecase.CreateOrder(r.Context(), &order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	utils.JSONResponse(w, http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

func (h *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.usecase.GetOrders(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.JSONResponse(w, http.StatusOK, orders)
}
