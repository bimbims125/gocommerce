package usecase

import (
	"fmt"
	"gocommerce/internal/entity"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type MidtransUsecase struct {
	Validate *validator.Validate
}

func NewMidtransUsecase(validate *validator.Validate) *MidtransUsecase {
	return &MidtransUsecase{Validate: validate}
}

func (mu *MidtransUsecase) CreatePayment(order entity.GetOrder, request entity.MidtransRequest) (entity.MidtransResponse, error) {
	if err := mu.Validate.Struct(request); err != nil {
		return entity.MidtransResponse{}, err
	}

	// Initialize snap client midtrans
	snapClient := snap.Client{}
	snapClient.New(os.Getenv("MIDTRANS_SERVER_KEY"), midtrans.Sandbox)

	// Create customer address
	custAddress := createCustomerAddress(order)

	// Create snap request
	req := createSnapRequest(order, request, custAddress)

	response, err := snapClient.CreateTransaction(req)
	if err != nil {
		return entity.MidtransResponse{}, err
	}

	midtransResponse := createMidtransResponse(response)
	return entity.MidtransResponse{
		TransactionID: (*req.Items)[0].ID,
		Token:         midtransResponse.Token,
		RedirectURL:   midtransResponse.RedirectURL,
	}, nil
}

func createCustomerAddress(order entity.GetOrder) *midtrans.CustomerAddress {
	return &midtrans.CustomerAddress{
		FName:       order.User.Name,
		Phone:       "083819588819",
		Address:     "Jl. Raya Bogor KM20",
		City:        "Bogor",
		Postcode:    "16911",
		CountryCode: "IDN",
	}
}
func createSnapRequest(order entity.GetOrder, request entity.MidtransRequest, custAddress *midtrans.CustomerAddress) *snap.Request {
	orderId := fmt.Sprintf("MIDTRANS%d%s%d", order.Product.ID, request.ItemID, time.Now().UnixNano())
	fmt.Println("Amount :", request.Amount)
	return &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderId,
			GrossAmt: request.Amount,
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName:    order.User.Name,
			Email:    order.User.Email,
			Phone:    "083819588819",
			BillAddr: custAddress,
			ShipAddr: custAddress,
		},
		EnabledPayments: snap.AllSnapPaymentType,
		Items: &[]midtrans.ItemDetails{
			{
				ID:    orderId,
				Qty:   int32(order.Quantity),
				Price: int64(order.Product.Price),
				Name:  order.Product.Name,
			},
		},
	}
}

func createMidtransResponse(response *snap.Response) entity.MidtransResponse {
	return entity.MidtransResponse{
		Token:       response.Token,
		RedirectURL: response.RedirectURL,
	}
}
