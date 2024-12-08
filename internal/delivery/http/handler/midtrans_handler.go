package handler

import (
	"encoding/json"
	"gocommerce/internal/entity"
	"gocommerce/internal/usecase"
	"io"
	"net/http"
)

type MidtransHandler struct {
	usecase      *usecase.MidtransUsecase
	orderUsecase *usecase.OrderUsecase
}

func NewMidtransHandler(usecase *usecase.MidtransUsecase, orderUseCase *usecase.OrderUsecase) *MidtransHandler {
	return &MidtransHandler{usecase: usecase, orderUsecase: orderUseCase}
}

func (h *MidtransHandler) PaymentCallbackHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		// utils.WriteError(w, http.StatusInternalServerError, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	callbackResponse := map[string]interface{}{}
	err = json.Unmarshal(body, &callbackResponse)
	if err != nil {
		// utils.WriteError(w, http.StatusInternalServerError, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var order entity.Order
	switch callbackResponse["transaction_status"].(string) {
	case "settlement":
		order.PaymentStatus = "completed"
	case "expire":
		order.PaymentStatus = "failed"
	case "cancel":
		order.PaymentStatus = "cancelled"
	default:
		order.PaymentStatus = "pending"
	}

	err = h.orderUsecase.UpdateOrderPaymentStatus(r.Context(), callbackResponse["order_id"].(string), order.PaymentStatus)
	if err != nil {
		// utils.WriteError(w, http.StatusInternalServerError, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return
}
