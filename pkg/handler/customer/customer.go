package customer

import (
	fs "stripe-example/external/firestore"
	st "stripe-example/external/stripe"

	"github.com/labstack/echo"
)

type Handler struct {
	stripe    *st.Stripe
	firestore *fs.FireStore
}

func NewCustomerHandler(stClient *st.Stripe, fsClient *fs.FireStore) *Handler {
	return &Handler{
		stripe:    stClient,
		firestore: fsClient,
	}
}

func (h *Handler) ListCustomer(c echo.Context) error {
	return nil

}
