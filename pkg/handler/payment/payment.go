package payment

import (
	fs "stripe-example/external/firestore"
	st "stripe-example/external/stripe"
)

type Handler struct {
	stripe    *st.Stripe
	firestore *fs.FireStore
}

func NewPaymentHandler(stClient *st.Stripe, fsClient *fs.FireStore) *Handler {
	return &Handler{
		stripe:    stClient,
		firestore: fsClient,
	}
}
