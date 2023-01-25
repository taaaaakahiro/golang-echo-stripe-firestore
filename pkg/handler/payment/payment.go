package payment

import (
	st "stripe-example/external/stripe"
)

type Handler struct {
	stripe *st.Stripe
}

func NewPaymentHandler(stClient *st.Stripe) *Handler {
	return &Handler{
		stripe: stClient,
	}
}
