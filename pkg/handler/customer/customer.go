package customer

import (
	st "stripe-example/external/stripe"

	"github.com/labstack/echo"
)

type Handler struct {
	stripe *st.Stripe
}

func NewCustomerHandler(stClient *st.Stripe) *Handler {
	return &Handler{
		stripe: stClient,
	}
}

func (h *Handler) ListCustomer(c echo.Context) error {
	return nil

}
