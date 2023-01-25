package handler

import (
	st "stripe-example/external/stripe"
	"stripe-example/pkg/handler/customer"
	"stripe-example/pkg/handler/payment"
	"stripe-example/pkg/handler/product"
)

type Handler struct {
	Customer *customer.Handler
	Payment  *payment.Handler
	Product  *product.Handler
}

func NewHandler(st *st.Stripe) (*Handler, error) {
	return &Handler{
		Customer: customer.NewCustomerHandler(st),
		Payment:  payment.NewPaymentHandler(st),
		Product:  product.NewProductHandler(st),
	}, nil
}
