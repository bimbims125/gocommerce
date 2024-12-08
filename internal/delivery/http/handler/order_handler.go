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
	usecase         *usecase.OrderUsecase
	midtransUsecase *usecase.MidtransUsecase
	userUsecase     *usecase.UserUseCase
	productUsecase  *usecase.ProductUseCase
}

func NewOrderHandler(usecase *usecase.OrderUsecase, midtransUsecase *usecase.MidtransUsecase, userUsecase *usecase.UserUseCase, productUsecase *usecase.ProductUseCase) *OrderHandler {
	return &OrderHandler{usecase: usecase, midtransUsecase: midtransUsecase, userUsecase: userUsecase, productUsecase: productUsecase}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {

	var order entity.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := h.productUsecase.GetProductByID(r.Context(), order.ProductID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.userUsecase.GetUserByID(r.Context(), order.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	getOrder := entity.GetOrder{
		User:     entity.UserOrder{ID: user.ID, Name: user.Name, Email: user.Email},
		Product:  entity.ProductOrder{ID: product.ID, Name: product.Name, Price: product.Price},
		Quantity: order.Quantity,
	}

	// Total amount
	totalAmount := getOrder.Product.Price * float64(order.Quantity)

	midtransResponse, err := h.midtransUsecase.CreatePayment(getOrder, entity.MidtransRequest{UserID: order.UserID, Amount: int64(totalAmount),
		ItemID: order.ProductID, ItemName: getOrder.Product.Name})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(totalAmount)
		return
	}

	// Transaction ID
	order.TransactionID = midtransResponse.TransactionID

	// Total Price
	order.TotalPrice = totalAmount

	id, err := h.usecase.CreateOrder(r.Context(), &order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	fmt.Print(id)
	utils.JSONResponse(w, http.StatusCreated, midtransResponse)
}

func (h *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.usecase.GetOrders(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.JSONResponse(w, http.StatusOK, orders)
}
