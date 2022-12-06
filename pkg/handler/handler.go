package handler

import (
	fs "stripe-example/external/firestore"
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

func NewHandler(st *st.Stripe, fs *fs.FireStore) *Handler {
	return &Handler{
		Customer: customer.NewCustomerHandler(st, fs),
		Payment:  payment.NewPaymentHandler(st, fs),
		Product:  product.NewProductHandler(st, fs),
	}
}
