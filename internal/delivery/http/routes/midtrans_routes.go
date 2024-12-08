package routes

import (
	"gocommerce/internal/delivery/http/handler"
	"gocommerce/internal/usecase"

	"github.com/gorilla/mux"
)

func RegisterMidtransRoutes(router *mux.Router, uc *usecase.MidtransUsecase, orderUsecase *usecase.OrderUsecase) {
	midtransHandler := handler.NewMidtransHandler(uc, orderUsecase)

	router.HandleFunc("/midtrans/callback", midtransHandler.PaymentCallbackHandler).Methods("POST")
}
